<template>
    <div v-if="loading">Loading</div>
    <scatter-chart
      v-else
      :data="listData"
      labels=""
      dataLabel="Electric Usage Data"
      xAxisLabel="watt"
      yAxisLabel="temp"
    />
  </template>
  
  <script lang="ts">
  import { getData } from "../function/api";
  import { onMounted, ref } from "vue";
  import ScatterChart from "@/components/charts/ScatterChart.vue";
  
  export default {
    name: "ElecTemp",
    components: {
      ScatterChart,
    },
    setup() {
        interface Point {
            x: number;
            y: number
        }
      const tmdData = ref([]);
      const elecData = ref([]);
      const listData = ref();
      const loading = ref(true);
  
      async function getRainData() {
        const ELECDATA = await getData(
          "https://daq.ku.sirateek.dev/api/elec/3hour"
        );
        const TMDDATA = await getData(
          "https://daq.ku.sirateek.dev/api/tmd/raw/data"
        );
        tmdData.value = TMDDATA.data;
        elecData.value = ELECDATA.data;
        loading.value = false;
        const scatterPoint : Point[] = []
        for(var i=0; i<tmdData.value.length;) {
            for(var j=0; j<elecData.value.length; j++) {
                if (elecData.value[j]["timestamp"] > tmdData.value[i]["timestamp"]) {
                    i++
                }
                var data = {x: elecData.value[j]["electric"], y:tmdData.value[i]["temp"]}
                scatterPoint.push(data)
            }
            break
        }
        listData.value = scatterPoint
      }

  
      onMounted(getRainData);
  
      return {
        tmdData,
        elecData,
        loading,
        listData
      };
    },
  };

  </script>
  