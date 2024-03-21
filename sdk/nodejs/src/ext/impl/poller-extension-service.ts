import {ExtensionService} from "../extension-service";
import {Error, Event} from "../../model";
import {Client} from "../../client";
import {AbstractExtensionService} from "./abstract-extension-service";
import {ExternalCall} from "../../model/extension";
import {ChannelEventPoller} from "../../impl/channel-event-poller";

export class PollerExtensionService extends AbstractExtensionService implements ExtensionService {
    private poller: ChannelEventPoller;

    constructor(serviceName: string, client: Client, private channelKey: string) {
        super(serviceName, client);
        this.poller = new ChannelEventPoller(client, channelKey, serviceName, this.handleEvent.bind(this));
    }

    handleEvent(event: Event): void {
        console.log("Received event: " + JSON.stringify(event));

        try {
            const processedEvent = this.processEvent(event);
            if (event.sync) {
                this.client.writeEvent(this.channelKey, processedEvent);
            }
        } catch (e: any) {
            console.error("Unable to process event", e);
            event.error = {
                message: e.message,
            } as Error;
            if (event.sync) {
                this.client.writeEvent(this.channelKey, event);
            }
        }
    }

    async run(): Promise<void> {
        await this.registerPendingItems()

        this.poller.start();
    }

    close(): void {
        this.poller.close()
    }

    protected prepareExternalCall(): ExternalCall {
        return {
            channelCall: {
                channelKey: this.channelKey,
            }
        } as ExternalCall;
    }
}
