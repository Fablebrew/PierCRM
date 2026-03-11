<template>
  <div class="chart-wrapper">
    <!-- Y-метки слева -->
    <div class="y-axis">
      <div v-for="label in yLabels" :key="label" class="y-label">{{ label }}</div>
    </div>

    <!-- Столбцы и линии -->
    <div class="chart-content">
      <!-- Столбцы -->
      <div class="chart-bars">
        <div
          v-for="(item, index) in chartData"
          :key="index"
          class="bar"
          :style="{ height: item.value + '%', backgroundColor: item.color || '#446BFA' }"
        ></div>
      </div>

      <!-- Линии расходов и доходов -->
      <svg class="chart-line" :width="svgWidth" :height="svgHeight">
        <!-- Красная линия: расходы -->
        <path
          ref="expensesPath"
          :d="smoothPath('expenses')"
          fill="none"
          stroke="#C80000"
          stroke-width="3"
          stroke-linecap="round"
        />
        <!-- Синяя линия: доходы -->
        <path
          ref="incomePath"
          :d="smoothPath('income')"
          fill="none"
          stroke="#3383FB"
          stroke-width="3"
          stroke-linecap="round"
        />
      </svg>

      <!-- X-метки -->
      <div class="x-axis">
        <div
          v-for="(item, index) in chartData"
          :key="index"
          class="x-label"
        >
          {{ item.label }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, nextTick } from 'vue';

export default {
  props: {
    chartData: {
      type: Array,
      required: true,
      // Пример данных:
      // [{ label: '1', value: 40, expenses: 40, income: 60 }, ...]
    },
    svgHeight: { type: Number, default: 400 },
  },
  setup(props) {
    const expensesPath = ref(null);
    const incomePath = ref(null);

    const svgWidth = props.chartData.length * 80;

    const yLabels = (() => {
      const maxValue = Math.max(
        ...props.chartData.flatMap(d => [d.expenses, d.income]),
        100
      );
      return [maxValue, maxValue * 0.75, maxValue * 0.5, maxValue * 0.25, 0];
    })();

    const smoothPath = type => {
      if (!props.chartData.length) return '';
      const h = props.svgHeight;
      const gap = 80;
      const points = props.chartData.map((d, i) => ({
        x: i * gap + gap / 2,
        y: h - (d[type] / 100) * h,
      }));

      let d = `M ${points[0].x},${points[0].y}`;
      for (let i = 0; i < points.length - 1; i++) {
        const cpX = (points[i].x + points[i + 1].x) / 2;
        d += ` C ${cpX},${points[i].y} ${cpX},${points[i + 1].y} ${points[i + 1].x},${points[i + 1].y}`;
      }
      return d;
    };

    const animatePath = pathEl => {
      const length = pathEl.getTotalLength();
      pathEl.style.strokeDasharray = length;
      pathEl.style.strokeDashoffset = length;
      pathEl.getBoundingClientRect(); // trigger reflow
      pathEl.style.transition = 'stroke-dashoffset 2s ease-in-out';
      pathEl.style.strokeDashoffset = '0';
    };

    onMounted(() => {
      nextTick(() => {
        if (expensesPath.value) animatePath(expensesPath.value);
        if (incomePath.value) animatePath(incomePath.value);
      });
    });

    return { expensesPath, incomePath, svgWidth, yLabels, smoothPath };
  },
};
</script>

<style scoped>
.chart-wrapper {
  display: flex;
  width: 100%;
  padding: 20px;
  box-sizing: border-box;
  overflow-x: auto;
}

.y-axis {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  margin-right: 10px;
  height: 400px;
}

.y-label {
  font-family: 'Manrope';
  font-size: 14px;
  color: rgba(0,0,0,0.5);
  text-align: right;
}

.chart-content {
  position: relative;
  flex: 1;
  min-width: 400px;
}

.chart-bars {
  display: flex;
  align-items: flex-end;
  gap: 20px;
  height: 100%;
}

.bar {
  width: 40px;
  border-radius: 6px;
  opacity: 0.6;
}

.chart-line {
  position: absolute;
  top: 0;
  left: 0;
  pointer-events: none;
}

.x-axis {
  display: flex;
  justify-content: space-between;
  margin-top: 10px;
}

.x-label {
  width: 40px;
  text-align: center;
  font-family: 'Manrope';
  font-size: 14px;
  color: rgba(0,0,0,0.5);
}
</style>