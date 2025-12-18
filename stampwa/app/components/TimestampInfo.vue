<script setup lang="ts">
import type { Uint256 } from 'web3';

const { chainId, contractAddress, hash, tree, parents, timestamp } = defineProps({
    chainId: Number,
    walletAddress: String,
    contractAddress: String,
    hash: String,
    tree: String,
    parents: Array as PropType<string[]>,
    timestamp: Object as PropType<BigInt>,
})

const network = await ChainNetworkInfo(chainId || 0)

let items = computed(() =>
    [
        {
            icon: 'i-lucide-link-2',
            title: 'Chainstamping Git Commit Info',
            description: '  ',
        },
        {
            icon: 'i-lucide-network',
            title: 'Network',
            description: `${network.name} - 0x${network.chainId.toString(16)}`
        },
        {
            icon: 'i-lucide-shield-check',
            title: 'Contract Address',
            description: contractAddress || 'No contract address provided'
        },
        {
            icon: 'i-lucide-hash',
            title: 'Commit Hash',
            description: hash || 'No commit hash provided'
        },
        {
            icon: 'i-lucide-folder',
            title: 'Tree Hash',
            description: tree || 'No tree hash provided'
        }
    ].concat(parents?.length ? parents.map(parent => ({
        icon: 'i-lucide-git-commit',
        title: 'Parent Commit Hash',
        description: parent
    })) : [{
        icon: 'i-lucide-git-commit',
        title: 'Parent Commit Hash',
        description: 'No parent commits provided'
    }]).concat({
        icon: 'i-lucide-timer',
        title: 'Timestamp',
        description: timestamp?.toString() || 'No timestamp available'
    })
)
</script>

<template>
    <UCard variant="subtle">
        <UTimeline :items :default-value="0" />
        <slot />
    </UCard>
</template>