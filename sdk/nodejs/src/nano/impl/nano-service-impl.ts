import {Client} from "../../client";
import {NanoService} from "../nano-service";
import {Code, CodeEntityInfo, Language} from "../model/code";
import {Repository} from "../../repository";

export class NanoServiceImpl implements NanoService {
    private repo: Repository<Code>;

    public constructor(private client: Client) {
        this.repo = client.repo<Code>(CodeEntityInfo)
    }

    async deployCode(code: Code, override?: boolean): Promise<Code> {
        if (override) {
            return this.repo.apply(code)
        } else {
            return this.repo.create(code)
        }
    }

    deploy(name: string, language: Language, source: string, override?: boolean): Promise<Code> {
        return this.deployCode({
            name: name,
            language: language,
            content: Buffer.from(source, "utf-8").toString("base64"),
        } as Code, override)
    }

    unDeploy(code: Code): Promise<Code> {
        return this.repo.delete(code.id)
    }


}