<template>
    <div v-if="loading">Loading</div>
    <line-chart
      v-else
      :data="waterData"
      :labels="tempData"
      dataLabel="Water Usage Data"
      xAxisLabel="temp"
      yAxisLabel="water"
    />
  </template>
  
  <script lang="ts">
  import { getData } from "../function/api";
  import { onMounted, ref } from "vue";
  import LineChart from "@/components/charts/LineChart.vue";
  
  export default {
    name: "WaterTemp",
    components: {
      LineChart,
    },
    setup() {
      const tempData = ref([]);
      const waterData = ref([]);
      const loading = ref(true);
  
      async function getRainData() {
        const WATERDATA = await getData(
          "https://daq.ku.sirateek.dev/api/water/perday"
        );
        const TMDDATA = await getData(
          "https://daq.ku.sirateek.dev/api/tmd/perday"
        );
        tempData.value = TMDDATA.data["temp"];
        waterData.value = WATERDATA.data;
        loading.value = false;
      }
  
      onMounted(getRainData);
  
      return {
        tempData,
        waterData,
        loading,
      };
    },
  };
  </script>
  