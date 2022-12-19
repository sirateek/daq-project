<template>
    <header class="sticky top-0 bg-lime-700 shadow-lg z-50">
        <nav class=" flex sm:flex-row items-center p-6 text-white ">
            <RouterLink :to="{name: 'home'}">
                <div class="flex items-center gap-3">
                    <i class="fa-sharp fa-solid fa-sun"></i>
                    <p class="text-2xl"> Happy E Water</p>
                </div>
            </RouterLink>
            <div class="flex gap-3 flex-1 justify-end">
                <i class="fa-solid fa-circle-info text-xl hover:text-black duration-150 cursor-pointer" @click="toggleInfo"></i>
            </div>
            <InfoModel :modalActive="modalActive" @close-modal="toggleInfo">
                <div class="text-black">
                    <h1 class="text-2xl mb-1">About:</h1>
                    <p class="mb-4">The Happy E Water is a project to find relation between electric, water, temperature and rainfall</p>
                    <h2 class="text-2xl">How it works:</h2>
                    <ol class="list-decimal list-inside mb-4">
                        <li>We will measure the electricity, water by using the SCT-013 and YF-S201.</li>
                        <li>Temperater and rainfall we will get data from the Thai Meteorological Department API.</li>
                        <li>We will convert data in to the graph to show the relationship of how temperature and rainfall 
                            effect &nbsp;&nbsp;&nbsp;&nbsp;  used of electricity and water usage.</li>
                    </ol>
                </div>
            </InfoModel>
        </nav>   
    </header>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { RouterLink } from 'vue-router';
import InfoModel from '@/components/InfoModel.vue';

const modalActive= ref(false);
const toggleInfo = () => {
    modalActive.value = !modalActive.value;
}
</script>


