<template>
  <div class="chart-container">
    <canvas ref="chart"></canvas>
  </div>
</template>

<script>
import {
  Chart,
  ArcElement,
  Tooltip,
  Legend,
  DoughnutController
} from "chart.js";

Chart.register(
  DoughnutController,
  ArcElement,
  Tooltip,
  Legend
);

export default {
  name: "DonutChart",

  mounted() {
    const ctx = this.$refs.chart;

    new Chart(ctx, {
      type: "doughnut",
      data: {
        labels: [
          "Операционные расходы",
          "Маркетинг",
          "Зарплаты",
          "Налоги",
          "Прочее"
        ],
        datasets: [
          {
            data: [9000, 5000, 8000, 4000, 3000],
            backgroundColor: [
              "#EC1A23",
              "#2065FA",
              "#801EE8",
              "#FF32A1",
              "#62EAFC"
            ],
            borderWidth: 0
          }
        ]
      },

      options: {
        cutout: "65%", // размер дырки
        plugins: {
          legend: {
            position: false
          }
        }
      }
    });
  }
};
</script>

<style>
.chart-container {
  width: 200px;
  height: 200px;
}
</style>