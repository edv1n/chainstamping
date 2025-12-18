import type { Contract, Uint256 } from 'web3';
import Web3 from 'web3';
import type { Chainstamper$Type } from './contracts/Chainstamper.sol/artifacts';
import metadata from './contracts/Chainstamper.sol/Chainstamper.json';
import { th } from '@nuxt/ui/runtime/locale/index.js';


export default function (web3: Web3 = new Web3((window as any).ethereum), contractAddress: string): Chainstamper {
    return new Chainstamper(web3, contractAddress);
}

class Chainstamper {
    // Chainstamper methods and properties would go here

    contractAddress: string = ''
    web3: Web3 = new Web3((window as any).ethereum)
    contract: Contract<Chainstamper$Type["abi"]>

    constructor(web3: Web3, contractAddress: string) {
        this.web3 = web3;

        if (contractAddress) {
            this.contractAddress = contractAddress;
        }

        this.contract = new this.web3.eth.Contract((metadata as Chainstamper$Type).abi, this.contractAddress);
    }

    async stampCommit(hash: string, tree: string, parents: string[]): Promise<void> {
        let tx = (() => {
            try {
                return this.contract.methods.stampCommit({ hash, tree, parents })
            } catch (error) {
                throw new Error(`Failed to create stampCommit transaction: ${error}`);
            }
        })()

        let accounts = await (async () => {
            try {
                return await this.web3.eth.getAccounts();
            } catch (error) {
                throw new Error(`Failed to get accounts: ${error}`);
            }
        })()

        let account = accounts[0] || (() => { throw new Error('No accounts available') })();

        let gas = await (async () => {
            try {
                return tx.estimateGas()
            } catch (error) {
                throw new Error(`Failed to estimate gas: ${error}`);
            }
        })()

        let receipt = await (async () => {
            try {
                return await tx.send({ from: account, gas: gas.toString() });
            } catch (error) {
                throw new Error(`Failed to send transaction: ${error}`);
            }
        })()
    }

    async getTimestamp(hash: string, tree: string, parents: string[]): Promise<Uint256> {
        try {
            let timestamp: Uint256 = await this.contract.methods.getTimestamp({ hash, tree, parents }).call();
            return timestamp;
        } catch (error) {
            let { innerError: { message } } = JSON.parse(JSON.stringify(error));
            if (message == "execution reverted: Commit not timestamped") {
                return "" as Uint256;
            }

            throw new Error(`Failed to get timestamp: ${error}`);
        }
    }
}