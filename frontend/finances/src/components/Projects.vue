<template>
  <Header />
  <Nav pageTitle="Проекты" />

  <section id="main">

    <!-- фильтры -->
    <div class="filters">
      <div class="row between">

        <div class="filters-row row">
          <button class="filter active">По умолчанию</button>
          <button class="filter">Этап</button>
          <button class="filter">Дата создания</button>
          <button class="filter">Дата изменения</button>
          <button class="filter">До дедлайна</button>
          <button class="filter">Исполнитель</button>
          <button class="filter">Период</button>
          <button class="filter" @click="showModal = true">
            <img src="@/assets/img/Vector.svg" alt="add_project">
          </button>
        </div>

        <div class="download-projects">
          <button>Скачать</button>
        </div>

      </div>
    </div>

    <!-- контейнер проектов -->
    <div class="projects-container">

      <!-- header -->
      <div class="projects-header grid">
        <div>Название проекта</div>
        <div>Приоритет</div>
        <div>Этап</div>
        <div>Начало</div>
        <div>До дедлайна</div>
        <div>Стоимость</div>
        <div>Оплачено</div>
        <div>Остаток</div>
        <div>Прочий доход</div>
        <div>Выручка</div>
        <div>Расход</div>
        <div class="profit">Прибыль</div>
      </div>

      <!-- строки -->
        <div
            v-for="(project, index) in projects"
            :key="index"
            class="project-row grid"
        >
            <div class="name">
                <input type="checkbox">
                {{ project.name }}
            </div>

            <div>{{ project.priority }}</div>
            <div>{{ project.stage }}</div>
            <div>{{ project.start }}</div>
            <div class="deadline">{{ project.deadline }}</div>
            <div>{{ project.cost }}</div>
            <div>{{ project.paid }}</div>
            <div>{{ project.remainder }}</div>
            <div>{{ project.otherIncome }}</div>

            <div>{{ project.revenue }}</div>
            <div>{{ project.expense }}</div>
            <div class="profit">{{ project.profit }}</div>

            <!-- кнопки -->
            <div class="finance-buttons">
                <button class="income" @click="openIncomeModal(project)">Доход</button>
                <button class="expense-btn" @click="openExpenseModal(project)">Расход</button>
            </div>
        </div>

    </div>
    <div class="bottom-button">
        <button><a href="#">Создать проект</a></button>
    </div>
  </section>
  <div v-if="showModal" class="modal-overlay" @click.self="showModal = false">

  <div class="modal">

    <div class="modal-header">
      <span>Создать проект</span>
      <button class="close" @click="showModal = false">✕</button>
    </div>

    <div class="modal-body">

      <div class="form-grid">

        <div class="field">
          <label>Название проекта</label>
          <input type="text">
        </div>

        <div class="field">
          <label>Приоритет</label>
          <select>
            <option>Низкий</option>
            <option>Средний</option>
            <option>Высокий</option>
          </select>
        </div>
        <div class="field">
          <label>Этап</label>
          <input type="text">
        </div>

        <div class="field">
          <label>Начало проекта</label>
          <input type="date">
        </div>

        <div class="field">
          <label>Конец проекта</label>
          <input type="date">
        </div>

        <div class="field">
          <label>Стоимость</label>
          <input type="number">
        </div>

        <div class="field">
          <label>Оплачено</label>
          <select>
            <option value="">Да</option>
            <option value="">Нет</option>
          </select>
        </div>

        <div class="field">
          <label>Заказчик</label>
          <input type="text">
        </div>

        <div class="field">
          <label>Исполнитель</label>
          <input type="text">
        </div>
        <div class="field">
          <label>Налоговая ставка</label>
          <select>
            <option value="">6%</option>
            <option value="">12%</option>
          </select>
        </div>
        <div class="field">
          <label>Счет</label>
          <select>
            <option value="">Карта тинька</option>
            <option value="">Сбер 1</option>
            <option value="">Сбер 2</option>
          </select>
        </div>

      </div>

      <textarea placeholder="Описание проекта"></textarea>

    </div>

    <div class="modal-footer">

      <button class="save">Сохранить</button>
      <button class="cancel" @click="showModal = false">Отменить</button>
      <button class="delete">Удалить данные</button>

    </div>

  </div>
</div>
<!-- INCOME MODAL -->
  <div v-if="showIncomeModal" class="modal-overlay" @click.self="showIncomeModal=false">

  <div class="modal">

    <div class="modal-header">
      <span>Добавить доход</span>
      <button class="close" @click="showIncomeModal=false">✕</button>
    </div>

    <div class="modal-body">

      <div class="form-grid">

        <div class="field">
          <label>Название проекта</label>
          <input type="text" :value="currentProject?.name" disabled>
        </div>

        <div class="field">
          <label>Счет</label>
          <select>
            <option>Не выбрано</option>
            <option>Карта тинька</option>
            <option>Сбер 1</option>
          </select>
        </div>

        <div class="field">
          <label>Сумма</label>
          <input type="number" placeholder="Не указано">
        </div>

        <div class="field">
          <label>Оплачено</label>
          <select>
            <option>Не указано</option>
            <option>Да</option>
            <option>Нет</option>
          </select>
        </div>

        <div class="field">
          <label>Дата</label>
          <input type="date">
        </div>

        <div class="field">
          <label>Ответственный</label>
          <input type="text" placeholder="Не указано">
        </div>

      </div>

      <textarea placeholder="Комментарий"></textarea>

    </div>

    <div class="modal-footer">
      <button class="save">Сохранить</button>
      <button class="cancel" @click="showIncomeModal=false">Отменить</button>
      <button class="delete">Удалить данные</button>
    </div>

  </div>

</div>
<!-- /INCOME MODAL -->
 <!-- EXPENSE MODAL -->
<div v-if="showExpenseModal" class="modal-overlay" @click.self="showExpenseModal=false">

  <div class="modal">

    <div class="modal-header">
      <span>Добавить расход</span>
      <button class="close" @click="showExpenseModal=false">✕</button>
    </div>

    <div class="modal-body">

      <div class="form-grid">

        <div class="field">
          <label>Название проекта</label>
          <input type="text" :value="currentProject?.name" disabled>
        </div>

        <div class="field">
          <label>Счет</label>
          <select>
            <option>Не выбрано</option>
            <option>Карта тинька</option>
            <option>Сбер 1</option>
          </select>
        </div>

        <div class="field">
          <label>Стоимость</label>
          <input type="number" placeholder="Не указано">
        </div>

        <div class="field">
          <label>Оплачено</label>
          <select>
            <option>Не указано</option>
            <option>Да</option>
            <option>Нет</option>
          </select>
        </div>

        <div class="field">
          <label>Дата</label>
          <input type="date">
        </div>

        <div class="field">
          <label>Исполнитель</label>
          <input type="text" placeholder="Не указано">
        </div>

      </div>

      <textarea placeholder="Комментарий"></textarea>

    </div>

    <div class="modal-footer">
      <button class="save">Сохранить</button>
      <button class="cancel" @click="showExpenseModal=false">Отменить</button>
      <button class="delete">Удалить данные</button>
    </div>

  </div>

</div>
<!-- /EXPENSE MODAL -->
</template>

<script>
import { ref } from "vue";
import Header from './Header.vue';
import Nav from './Nav.vue';

export default {
  components: {
    Header,
    Nav,
  },
  name: "Projects",
  data() {
    return {
      showModal: false,
      showIncomeModal: false,
      showExpenseModal: false,
      currentProject: null,
      projects: [
        { name: "Журнал Эксперт", priority: "Высокий", stage: "Дизайн", start: "30.07.2024", deadline: "51 день", cost: "35.000 р.", paid: "25.000 р.", remainder: "10.000 р.", otherIncome: "10.000 р.", revenue: "45.000 р.", expense: "5.000 р.", profit: "40.000 р." },
        { name: "Дизайн сайта", priority: "Средний", stage: "Верстка", start: "21.07.2024", deadline: "55 дней", cost: "75.000 р.", paid: "25.000 р.", remainder: "10.000 р.", otherIncome: "10.000 р.", revenue: "45.000 р.", expense: "5.000 р.", profit: "40.000 р." },
        { name: "Дизайн приложения", priority: "Низкий", stage: "Дизайн", start: "14.08.2024", deadline: "81 день", cost: "95.600 р.", paid: "25.000 р.", remainder: "10.000 р.", otherIncome: "10.000 р.", revenue: "45.000 р.", expense: "5.000 р.", profit: "40.000 р." },
      ]
    };
  },
  methods: {
    openIncomeModal(project) {
      this.currentProject = project
      this.showIncomeModal = true
    },
    openExpenseModal(project) {   // ← новый метод
      this.currentProject = project
      this.showExpenseModal = true
    },
  },
};

// Примеры подсветки (полупрозрачные фоны)
const highlights = ref([
  { id: 1, top: 255, height: 118 },
]);
</script>

<style scoped>

/* overlay */

.modal-overlay {
  position: fixed;
  inset: 0;

  background: rgba(0,0,0,0.35);
  backdrop-filter: blur(6px);

  display: flex;
  justify-content: center;
  align-items: center;

  z-index: 1000;
}

/* окно */

.modal {
  width: 755px;
  background: white;
  border-radius: 10px;
  overflow: hidden;
}

/* header */

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;

  padding: 18px 25px;
  border-bottom: 1px solid #BBBBBB;

  font-size: 20px;
}

.close {
  background: none;
  border: none;
  font-size: 20px;
  cursor: pointer;
}

/* body */

.modal-body {
  padding: 25px;
}

.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
}

.field {
  display: flex;
  flex-direction: column;
}

.field label {
  font-size: 14px;
  margin-bottom: 6px;
}

.field input,
.field select {
  height: 40px;
  border-radius: 6px;
  border: none;
  background: #F1F4FF;
  padding: 0 10px;
}

/* textarea */

textarea {
  width: 100%;
  margin-top: 20px;
  height: 120px;
  border-radius: 6px;
  border: none;
  background: #F1F4FF;
  padding: 10px;
  resize: none;
}

/* footer */

.modal-footer {
  display: flex;
  gap: 10px;
  padding: 20px 25px;
}

.save {
  background: #3383FB;
  color: white;
  border: none;
  padding: 8px 18px;
  border-radius: 4px;
}

.cancel {
  background: #E9E9E9;
  border: none;
  padding: 8px 18px;
  border-radius: 4px;
}

.delete {
  border: 1px solid #3383FB;
  background: white;
  padding: 8px 18px;
  border-radius: 4px;
}

template {
    height: 100vh;
}
.bottom-button {
    width: 100%;
    margin-top: 1%;
}
.bottom-button button {
    font-family: "Manrope";
    font-weight: 400;
    font-style: "Regular";
    font-size: 22px;
    line-height: 150%;
    letter-spacing: 0%;
    background-color:#446BFA;
    color: #FFFFFF;
    outline: none;
    border: none;
    border-radius: 16px;
    width: 100%;
    padding: 20px;
}
.bottom-button button a {
    text-decoration: none;
    color: #FFFFFF;
}
.bottom-button button:hover {
    cursor: pointer;
}
#main {
  padding: 20px;
  background: #F1F4FF;
}

/* контейнер */

.projects-container {
  margin-top: 40px;
  background: white;
  border: 1px solid #BBBBBB;
  border-radius: 15px;
  overflow: hidden;
}

/* grid колонок */

.grid {
  display: grid;

  grid-template-columns:
  minmax(180px, 2fr)   /* название */
  minmax(80px, 1fr)
  minmax(80px, 1fr)
  minmax(90px, 1fr)
  minmax(100px, 1fr)
  minmax(100px, 1fr)
  minmax(100px, 1fr)
  minmax(100px, 1fr)
  minmax(110px, 1fr)
  minmax(100px, 1fr)
  minmax(90px, 1fr)
  minmax(100px, 1fr);

  align-items: center;
}

/* header */

.projects-header {
  background: #DBE1F6;
  font-weight: 500;
  font-size: 14px;
  padding: 15px 20px;
  border-bottom: 1px solid #BBBBBB;
}

/* строка */

.project-row {
  padding: 15px 20px;
  border-bottom: 1px solid #BBBBBB;
  font-size: 14px;
  color: rgba(0,0,0,0.7);
}
.finance-buttons {
  grid-column: 10 / 13; /* колонки Выручка - Прибыль */

  display: flex;
  gap: 10px;
  margin-top: 10px;
}
.income {
  flex: 1;
  padding: 8px 0;
  background: #2EB352;
  border: none;
  border-radius: 4px;
  color: white;
  font-size: 12px;
  cursor: pointer;
}
.expense-btn {
  flex: 1;
  padding: 8px 0;
  background: #EC1A23;
  border: none;
  border-radius: 4px;
  color: white;
  font-size: 12px;
  cursor: pointer;
}

.income:hover {
  background: #27a049;
}

.expense-btn:hover {
  background: #d0171f;
}

.project-row:last-child {
  border-bottom: none;
}

/* название */

.name {
  display: flex;
  align-items: center;
  gap: 10px;
}

/* дедлайн */

.deadline {
  color: #3383FB;
}

/* прибыль */

.profit {
  font-weight: 600;
}

/* фильтры */

.row {
  display: flex;
  align-items: center;
}

.between {
  justify-content: space-between;
}

.filters-row {
  gap: 10px;
}

.filter {
  padding: 4px 10px;
  font-size: 10px;
  background: white;
  border: 1px solid #d8d8d8;
  border-radius: 4px;
  cursor: pointer;
}

.filter.active {
  background: #3383FB;
  color: white;
}

.download-projects button {
  padding: 5px 12px;
  font-size: 12px;
  border-radius: 4px;
  border: none;
  background: #3383FB;
  color: white;
}
</style>