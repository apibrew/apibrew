
import { Client } from '../client';
import { Entity } from "../entity";
import { Repository } from "../repository";
import { EntityInfo } from "../entity-info";
import { Container } from "../container";
import {BooleanExpression, Code, Event} from "../model/extension";
import axios from "axios/index";
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
                console.log("Begin Polling channel: " + this.channelKey);

                const result = await axios.get(Urls.eventsUrl(this.client.getUrl()) + "?channelKey=" + this.channelKey, {
                    headers: this.client.headers(),
                    validateStatus: (status) => true,
                    responseType: "stream",
                    timeout: 10 * 1000,
                });

                if (!this.isRunning()) {
                    return;
                }

                if (result.status !== 200) {
                    throw new ApiException(Code.INTERNAL_ERROR, result.statusText);
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
                    setTimeout(() => this.run(), 1000);
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
