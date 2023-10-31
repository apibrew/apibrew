import {Code, Language} from "./model/code";

export interface NanoService {
    deployCode(code: Code, override?: boolean): Promise<Code>;

    deploy(name: string, language: Language, source: string, override?: boolean): Promise<Code>;

    unDeploy(code: Code): Promise<Code>;
}
