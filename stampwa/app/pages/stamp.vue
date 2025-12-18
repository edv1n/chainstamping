<script setup lang="ts">
import { Web3 } from 'web3';

const route = useRoute()
const router = useRouter()

const title = 'Stamp Git Commit on Ethereum'
const description = 'Easily timestamp your Git commits on the Ethereum blockchain with Chainstamping.'

let chainId = Number(route.query.chain) || 11155111 // Sepolia Testnet
let contractAddress = route.query.contract as string
let hash = route.query.hash as string || ''
let tree = route.query.tree as string || ''
let queryParents = route.query.parent
let parents: string[] = queryParents
    ? Array.isArray(queryParents)
        ? queryParents as string[]
        : [queryParents as string]
    : []


let walletAddress = ref('')
let web3 = await (async () => {
    try {
        let network = await ChainNetworkInfo(chainId)

        return new Web3(network.rpc)
    } catch (e) {
        console.error('Error initializing web3 without wallet:', e)
    }
})();

(web3) && (await (async () => {
    try {
        console.log('Checking for existing timestamp...')
        if (await Chainstamper(web3, contractAddress).getTimestamp(hash, tree, parents)) {
            router.push({
                name: 'info',
                query: {
                    chain: chainId,
                    contract: contractAddress,
                    hash,
                    tree,
                    parent: parents
                }
            })
        }
    } catch (e) {
        console.error('Error checking existing timestamp:', e)
    }
})());

(async () => {
    try {
        web3 = await Web3WalletEnabled(chainId)
        if (!web3) {
            return
        }

        let walletAccounts = await web3.eth.getAccounts()
        let account = walletAccounts[0] || ''
        console.log('Connected wallet account:', account)
        walletAddress.value = account
    } catch (e) {
        console.error('Error connecting to wallet:', e)
    }
})()

let click = async () => {
    await Chainstamper(web3, contractAddress).stampCommit(hash, tree, parents)
    router.push({
        name: 'info',
        query: {
            chain: chainId,
            contract: contractAddress,
            hash,
            tree,
            parent: parents
        }
    })
}
</script>

<template>
    <StampingInfo :chain-id="chainId" :contract-address :walletAddress :hash="hash" :tree="tree" :parents="parents">
        <UButton v-is="walletAddress" :disabled="!walletAddress" @click="click">
            Stamp Commit
        </UButton>
    </StampingInfo>
</template>