import {Client} from '../client';
import {BooleanExpression, Code, Event} from "../model/extension";
import axios from "axios";
import {Urls} from "./client-impl";
import {ApiException} from "../api-exception";


export class ChannelEventPoller {

    constructor(
        private client: Client,
        private channelKey: string,
        private threadName: string,
        private consumer: (event: Event) => void,
    ) {
    }

    private stillRunning: boolean = false;
    private isStopped: boolean = false;

    public start() {
        this.ensureCanRun();
        setTimeout(() => this.run(), 0);
    }

    private ensureCanRun() {
        if (this.isStopped) {
            throw new Error("Poller is stopped");
        }

        if (this.stillRunning) {
            throw new Error("Poller is already running");
        }
    }

    private async run() {
        this.stillRunning = true;

        while (this.isRunning()) {
            try {
                const url = Urls.eventsUrl(this.client.getUrl()) + "?channelKey=" + this.channelKey

                console.log("Begin Polling channel: " + this.channelKey + " at: " + url);

                const result = await axios.get(url, {
                    headers: this.client.headers(),
                    validateStatus: (status) => true,
                    responseType: "stream",
                    timeout: 10 * 1000,
                });

                if (!this.isRunning()) {
                    return;
                }

                if (result.status !== 200) {
                    console.log("Polling channel: " + this.channelKey + " failed: " + result.status + ": " + result.statusText);
                    await new Promise(resolve => setTimeout(resolve, 1000));
                    continue;
                }

                const stream = result.data;

                // block until stream is closed

                let wc = 0;

                wc++;

                stream.on('data', (data: Buffer) => {
                    const event = JSON.parse(data.toString()) as Event;

                    if (event.id === "heartbeat-message") {
                        console.log("Received heartbeat message");
                    } else {
                        console.log("Received event: " + event);
                        this.consumer(event);
                    }
                });

                stream.on('end', () => {
                    console.log("stream done");
                    wc--;
                });

                stream.on('error', (err: any) => {
                    console.log("stream error: " + err);
                    wc--;
                });

                while (wc > 0) {
                    await new Promise(resolve => setTimeout(resolve, 100));
                }

                console.log("Polling channel: " + this.channelKey + " complete");
            } catch (e) {
                console.log("Error polling channel: " + this.channelKey + ": " + e);
                try {
                    await new Promise(resolve => setTimeout(resolve, 1000));
                } catch (ex: any) {
                    throw new Error(ex.message);
                }
            }
        }
    }

    private isRunning() {
        return !this.isStopped && this.stillRunning;
    }

    public close() {
        this.isStopped = true;
        this.stillRunning = false;
    }

}
