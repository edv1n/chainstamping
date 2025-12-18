<script setup lang="ts">
import Web3 from 'web3';
import type TimestampInfoVue from '~/components/TimestampInfo.vue';

const route = useRoute()

const title = 'Stamp Git Commit on Ethereum'
const description = 'Easily timestamp your Git commits on the Ethereum blockchain with Chainstamping.'

let chainId = Number(route.query.chain) || 11155111 // Sepolia Testnet
let contractAddress = route.query.contract as string || '0xCaFF7E83bFCE9C9b968d79c500A6e78D34422B59'
let hash = route.query.hash as string || ''
let tree = route.query.tree as string || ''
let queryParents = route.query.parent
let parents: string[] = queryParents
    ? Array.isArray(queryParents)
        ? queryParents as string[]
        : [queryParents as string]
    : []

let timestamp = await (async () => {
    try {
        let network = await ChainNetworkInfo(chainId)
        let web3 = new Web3(network.rpc)
        return await Chainstamper(web3).getTimestamp(hash, tree, parents)
    } catch (e) {
        console.error('Error fetching timestamp:', e)
        return ""
    }
})()
</script>

<template>
    <TimestampInfo :chain-id="chainId" :contract-address="contractAddress" :hash="hash" :tree="tree" :parents="parents"
        :timestamp="timestamp">
    </TimestampInfo>
</template>