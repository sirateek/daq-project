<template>
  <div v-if="loading">Loading</div>
  <div v-else>
    <line-chart :data="elecData" :labels="rainData" dataLabel="Water Usage Data" xAxisLabel="rain" yAxisLabel="water"/>
  </div>
</template>
  
<script lang="ts">
import { getData } from "../function/api";
import { onMounted, ref } from "vue";
import LineChart from "@/components/charts/LineChart.vue";

export default {
  name: "WaterRain",
  components: {
    LineChart,
  },
  setup() {
    const rainData = ref([]);
    const elecData = ref([]);
    const loading = ref(true);

    async function getRainData() {
      const ELECDATA = await getData(
        "https://daq.ku.sirateek.dev/api/water/perday"
      );
      const TMDDATA = await getData(
        "https://daq.ku.sirateek.dev/api/tmd/perday"
      );
      rainData.value = TMDDATA.data["rain"];
      elecData.value = ELECDATA.data;
      loading.value = false;
    }

    onMounted(getRainData);

    return {
      rainData,
      elecData,
      loading,
    };
  },
};
</script>
  