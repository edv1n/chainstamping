export default async (chainId: number) => {
    try {
        let resp = await fetch('https://chainid.network/chains_mini.json')
        let networks: Networks = await resp.json();
        let network = networks.find(n => n.chainId === chainId);
        if (!network) {
            throw new Error(`Network with chainId ${chainId} not found`);
        }

        return {
            name: network.name,
            chainId: network.chainId,
            rpc: network.rpc.filter(url => url.includes('.publicnode.com'))[0] || '',
        }
    } catch (error) {
        return {
            name: 'Unknown Network',
            chainId,
            rpc: '',
        }
    }
}

type Networks = [{
    name: string;
    chainId: number;
    rpc: string[];
}]