import {Watcher} from "../watcher";
import {Entity} from "../entity";
import {EntityInfo} from "../entity-info";
import {Client} from "../client";
import {Code, Event} from "../model/extension";
import {ApiException} from "../api-exception";
import axios from "axios";
import {Urls} from "./client-impl";


export class WatcherImpl<T extends Entity> implements Watcher<T> {
    stillRunning: boolean = false;
    isStopped: boolean = false;

    constructor(private client: Client, private entityInfo: EntityInfo, private consumer: (event: Event) => void) {

    }

    public start() {
        this.ensureCanRun();
        this.stillRunning = true;
        this.run();
    }

    public async run() {
        this.ensureCanRun();
        this.stillRunning = true;

        while (this.isRunning()) {
            try {
                console.log("Begin watching: ", this.entityInfo);

                const result = await axios.get(Urls.recordWatchUrl(this.client.getUrl(), this.entityInfo.restPath), {
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
                })

                while (wc > 0) {
                    await new Promise(resolve => setTimeout(resolve, 100));
                }

                console.log("watch completed: " + this.entityInfo);
            } catch (e) {
                console.log("Error watching: " + e);
                try {
                    await new Promise(resolve => setTimeout(resolve, 1000));
                } catch (ex: any) {
                    throw new ApiException(Code.INTERNAL_ERROR, ex.message);
                }
            }
        }
    }

    private ensureCanRun() {
        if (this.isStopped) {
            throw new ApiException(Code.INTERNAL_ERROR, "Poller is stopped");
        }

        if (this.stillRunning) {
            throw new ApiException(Code.INTERNAL_ERROR, "Poller is already running");
        }
    }

    private isRunning() {
        return !this.isStopped && this.stillRunning;
    }

    stop(): void {
        this.isStopped = true;
        this.stillRunning = false;
    }
}
