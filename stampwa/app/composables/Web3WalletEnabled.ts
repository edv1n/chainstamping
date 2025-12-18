import Web3 from 'web3';

declare global {
    interface Window {
        ethereum?: any;
    }
}

export default async (chainId: number) => {
    if (!(window as any).ethereum) {
        throw new Error('No Ethereum provider found. Please install MetaMask.');
    }

    // Create a new Web3 instance using the provider
    const web3 = new Web3(window.ethereum as any);

    await window.ethereum?.request({
        method: 'wallet_switchEthereumChain',
        params: [{ chainId: `0x${chainId.toString(16)}` }],
    });

    await window.ethereum?.request({ method: 'eth_requestAccounts' });

    return web3;
}