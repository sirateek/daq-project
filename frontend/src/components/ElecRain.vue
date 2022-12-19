<template>
  <div v-if="loading">Loading</div>
  <line-chart
    v-else
    :data="elecData"
    :labels="rainData"
    dataLabel="Electric Usage Data"
    xAxisLabel="rain"
    yAxisLabel="watt"
  />
</template>

<script lang="ts">
import { getData } from "../function/api";
import { onMounted, ref } from "vue";
import LineChart from "@/components/charts/LineChart.vue";

export default {
  name: "ElecRain",
  components: {
    LineChart,
  },
  setup() {
    const rainData = ref();
    const elecData = ref();
    const loading = ref(true);

    async function getRainData() {
      const ELECDATA = await getData(
        "https://daq.ku.sirateek.dev/api/elec/3hour"
      );
      const TMDDATA = await getData(
        "https://daq.ku.sirateek.dev/api/tmd/raw/data"
      );
      rainData.value = TMDDATA.data;
      elecData.value = ELECDATA.data;
      loading.value = false;
      const listEelect = [];
      const listRain = [];
      for (let i=0; i<rainData.value.length; i++){
        listRain.push(rainData.value[i]["rain"])
      }
      for (let i=0; i<elecData.value.length; i++){
        listEelect.push(elecData.value[i]["electric"])
      }
      rainData.value = listRain;
      elecData.value = listEelect;
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
