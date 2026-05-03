<template>
  <Header />
  <Nav page-title="Проекты" class="nav-component" />

  <section id="main">

    <!-- Заголовок страницы -->
    <h1 class="page-title">Проекты</h1>

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
          <button class="filter filter-add">
            <img src="@/assets/img/Vector.svg" alt="add_project">
          </button>
        </div>

        <div class="download-projects">
          <button>Скачать</button>
        </div>

      </div>
    </div>

    <!-- Поиск для мобильных -->
    <div class="mobile-search">
      <span class="search-icon-mob">🔍</span>
      <input type="text" placeholder="Поиск" class="search-input-mob">
    </div>

    <!-- Кнопка для мобильных -->
    <button class="mobile-add-btn" @click="showModal = true">
      <span>+</span>
    </button>

    <!-- контейнер проектов (десктоп) -->
    <div class="projects-container desktop-table">

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

    <!-- Мобильные карточки проектов -->
    <div class="mobile-projects">
      <div
        v-for="(project, index) in projects"
        :key="'mob-' + index"
        class="mobile-project-card"
      >
        <div class="mobile-project-header">
          <h3 class="mobile-project-name">{{ project.name }}</h3>
          <div class="mobile-project-deadline">
            <span class="calendar-icon">📅</span>
            <span>{{ project.deadline }}</span>
          </div>
        </div>

        <div class="mobile-project-stage">
          <span class="stage-label">{{ project.stage }}</span>
          <div class="stage-progress">
            <span class="progress-dot" v-for="i in 5" :key="i" :class="{ active: i <= 3 }"></span>
          </div>
        </div>

        <div class="mobile-project-finance">
          <div class="finance-item">
            <span class="finance-label">Стоимость</span>
            <span class="finance-value">{{ project.cost }}</span>
          </div>
          <div class="finance-item">
            <span class="finance-label">Оплачено</span>
            <span class="finance-value">{{ project.paid }}</span>
          </div>
          <div class="finance-item">
            <span class="finance-label">Остаток</span>
            <span class="finance-value">{{ project.remainder }}</span>
          </div>
          <div class="finance-item profit-item">
            <span class="finance-label">Прибыль</span>
            <span class="finance-value profit-value">{{ project.profit }}</span>
          </div>
        </div>

        <div class="mobile-project-buttons">
          <button class="income" @click="openIncomeModal(project)">Доход</button>
          <button class="expense-btn" @click="openExpenseModal(project)">Расход</button>
        </div>
      </div>
    </div>

    <div class="bottom-button desktop-btn">
      <button @click="showModal = true"><a href="#">Создать проект</a></button>
    </div>
  </section>

  <!-- Модальное окно создания проекта -->
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

  <!-- Модальное окно дохода -->
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

  <!-- Модальное окно расхода -->
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
</template>

<script>
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
    openExpenseModal(project) {
      this.currentProject = project
      this.showExpenseModal = true
    },
  },
};
</script>

<style scoped>
/* Базовые стили */
* {
  box-sizing: border-box;
}

/* Заголовок страницы */
.page-title {
  text-align: center;
  font-family: "Manrope";
  font-weight: 600;
  font-size: 18px;
  line-height: 140%;
  color: rgba(0, 0, 0, 0.7);
  margin-bottom: 20px;
  display: none;
}

/* Мобильная кнопка + */
.mobile-add-btn {
  display: none;
  position: absolute;
  top: 70px;
  right: 20px;
  width: 25px;
  height: 25px;
  border-radius: 4px;
  background: #3383FB;
  color: white;
  border: none;
  font-size: 25px;
  cursor: pointer;
  z-index: 10;
  align-items: center;
  justify-content: center;
  line-height: 1;
}

.mobile-add-btn span {
  line-height: 1;
}

/* Поиск для мобильных */
.mobile-search {
  display: none;
  position: relative;
  margin: 0 20px 20px;
}

.search-icon-mob {
  position: absolute;
  left: 12px;
  top: 50%;
  transform: translateY(-50%);
  font-size: 14px;
  z-index: 1;
}

.search-input-mob {
  width: 100%;
  height: 40px;
  border-radius: 29px;
  border: 0.5px solid #BBBBBB;
  background: #FFFFFF;
  padding: 0 16px 0 40px;
  font-family: 'Onest';
  font-size: 12px;
  color: #696969;
}

/* Мобильные карточки проектов */
.mobile-projects {
  display: none;
  flex-direction: column;
  gap: 0;
  margin-top: 20px;
}

.mobile-project-card {
  background: #FFFFFF;
  padding: 15px 20px;
  border-bottom: 1px solid #D7D7D7;
}

.mobile-project-card:last-child {
  border-bottom: none;
}

.mobile-project-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.mobile-project-name {
  font-family: 'Manrope';
  font-weight: 700;
  font-size: 14px;
  line-height: 140%;
  color: rgba(0, 0, 0, 0.7);
  margin: 0;
}

.mobile-project-deadline {
  display: flex;
  align-items: center;
  gap: 5px;
}

.calendar-icon {
  font-size: 14px;
}

.mobile-project-deadline span {
  font-family: 'Manrope';
  font-weight: 300;
  font-size: 14px;
  color: #494A4D;
}

.mobile-project-stage {
  display: flex;
  flex-direction: column;
  gap: 3px;
  margin-bottom: 15px;
}

.stage-label {
  font-family: 'Manrope';
  font-weight: 300;
  font-size: 9px;
  color: rgba(0, 0, 0, 0.7);
}

.stage-progress {
  display: flex;
  gap: 2px;
}

.progress-dot {
  width: 13px;
  height: 13px;
  background: #E8EAF2;
  border-radius: 1px;
}

.progress-dot.active {
  background: #8DCD94;
}

.mobile-project-finance {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
  margin-bottom: 15px;
}

.finance-item {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.finance-label {
  font-family: 'Manrope';
  font-weight: 300;
  font-size: 9px;
  color: rgba(0, 0, 0, 0.7);
}

.finance-value {
  font-family: 'Manrope';
  font-weight: 500;
  font-size: 14px;
  color: rgba(0, 0, 0, 0.7);
}

.profit-item .finance-value {
  color: #3383FB;
}

.mobile-project-buttons {
  display: flex;
  gap: 10px;
}

.mobile-project-buttons .income,
.mobile-project-buttons .expense-btn {
  flex: 1;
  padding: 9px 0;
  border: none;
  border-radius: 4px;
  color: white;
  font-family: 'Manrope';
  font-weight: 400;
  font-size: 12px;
  cursor: pointer;
  text-align: center;
}

/* Основные стили (десктоп) */
#main {
  padding: 20px;
  background: #F1F4FF;
  min-height: 100vh;
  position: relative;
}

.projects-container {
  margin-top: 40px;
  background: white;
  border: 1px solid #BBBBBB;
  border-radius: 15px;
  overflow: hidden;
}

.grid {
  display: grid;
  grid-template-columns:
    minmax(180px, 2fr)
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

.projects-header {
  background: #DBE1F6;
  font-weight: 500;
  font-size: 14px;
  padding: 15px 20px;
  border-bottom: 1px solid #BBBBBB;
}

.project-row {
  padding: 15px 20px;
  border-bottom: 1px solid #BBBBBB;
  font-size: 14px;
  color: rgba(0,0,0,0.7);
}

.project-row:last-child {
  border-bottom: none;
}

.name {
  display: flex;
  align-items: center;
  gap: 10px;
}

.deadline {
  color: #3383FB;
}

.profit {
  font-weight: 600;
}

.finance-buttons {
  grid-column: 10 / 13;
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

/* Фильтры */
.row {
  display: flex;
  align-items: center;
}

.between {
  justify-content: space-between;
}

.filters-row {
  gap: 10px;
  flex-wrap: wrap;
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

.bottom-button {
  width: 100%;
  margin-top: 1%;
}

.bottom-button button {
  font-family: "Manrope";
  font-weight: 400;
  font-size: 22px;
  line-height: 150%;
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

/* Модальные окна */
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.35);
  backdrop-filter: blur(6px);
  display: flex;
  justify-content: center;
  align-items: flex-start;
  z-index: 1000;
  padding: 20px;
  overflow-y: auto;
}

.modal {
  width: 100%;
  max-width: 755px;
  background: white;
  border-radius: 10px;
  margin: auto;
}

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

/* Адаптивность */
@media (max-width: 768px) {
  .nav-component {
    display: none !important;
  }

  .page-title {
    display: block;
  }

  #main {
    padding: 15px 0;
  }

  .filters {
    padding: 0 20px;
  }

  .filters-row {
    gap: 5px;
  }

  .filter {
    font-size: 9px;
    padding: 3px 8px;
  }

  .filter-add {
    display: none;
  }

  .download-projects {
    display: none;
  }

  .mobile-search {
    display: block;
  }

  .mobile-add-btn {
    display: flex;
  }

  .desktop-table {
    display: none;
  }

  .desktop-btn {
    display: none;
  }

  .mobile-projects {
    display: flex;
  }

  /* Адаптация модальных окон */
  .modal-overlay {
    padding: 10px;
    align-items: flex-start;
  }

  .modal {
    width: 100%;
    max-width: 100%;
    margin: 10px 0;
    border-radius: 10px;
  }

  .modal-header {
    padding: 14px 20px;
    font-size: 16px;
  }

  .modal-body {
    padding: 20px;
  }

  .form-grid {
    grid-template-columns: 1fr;
    gap: 15px;
  }

  .field label {
    font-size: 13px;
  }

  .field input,
  .field select {
    height: 38px;
    font-size: 14px;
  }

  textarea {
    height: 100px;
    font-size: 14px;
  }

  .modal-footer {
    flex-wrap: wrap;
    padding: 15px 20px;
    gap: 8px;
  }

  .save,
  .cancel,
  .delete {
    flex: 1;
    min-width: 80px;
    padding: 10px 12px;
    font-size: 13px;
    text-align: center;
  }
}

@media (max-width: 480px) {
  .page-title {
    font-size: 16px;
  }

  .filters {
    padding: 0 10px;
  }

  .mobile-search {
    margin: 0 10px 15px;
  }

  .mobile-add-btn {
    top: 60px;
    right: 10px;
  }

  .mobile-project-card {
    padding: 12px 15px;
  }

  .mobile-project-name {
    font-size: 12px;
  }

  .mobile-project-deadline span {
    font-size: 12px;
  }

  .finance-value {
    font-size: 12px;
  }

  /* Адаптация модальных окон для маленьких экранов */
  .modal-overlay {
    padding: 5px;
  }

  .modal-header {
    padding: 12px 15px;
    font-size: 15px;
  }

  .modal-body {
    padding: 15px;
  }

  .form-grid {
    gap: 12px;
  }

  .field label {
    font-size: 12px;
  }

  .field input,
  .field select {
    height: 36px;
    font-size: 13px;
  }

  textarea {
    height: 80px;
  }

  .modal-footer {
    padding: 12px 15px;
    flex-direction: column;
  }

  .save,
  .cancel,
  .delete {
    width: 100%;
    min-width: auto;
  }
}
</style>