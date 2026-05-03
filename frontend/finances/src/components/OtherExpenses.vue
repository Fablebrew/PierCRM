<template>
  <Header />
  <Nav page-title="Прочие расходы" class="nav-component"/>

  <section class="expenses-page">

    <!-- Заголовок страницы -->
    <h1 class="page-title">Прочие расходы</h1>

    <div class="expenses-container">

      <!-- Кнопка для мобильных -->
      <button class="mobile-add-btn" @click="showModal = true">
        <span>+</span>
      </button>

      <!-- header -->
      <div class="table-header table-grid">
        <div>Расход</div>
        <div>Сумма</div>
        <div>Дата оплаты</div>
        <div>Теги</div>
        <div></div>
      </div>

      <!-- строки -->
      <div
        v-for="(expense,index) in expenses"
        :key="index"
        class="table-row table-grid"
      >
        <div class="name">
          <input type="checkbox">
          <span class="expense-name">{{ expense.name }}</span>
        </div>

        <div class="expense-amount">{{ expense.amount }}</div>
        <div class="expense-date">{{ expense.date }}</div>

        <div class="tag">
            <span 
                class="tag-color"
                :style="{background: expense.color}"
            ></span>

            <span class="tag-text">
                {{ expense.tag }}
            </span>
        </div>

        <div class="edit">
          <img src="@/assets/img/redact.svg" class="edit-icon">
        </div>
      </div>
    </div>

    <!-- Кнопка для десктопа -->
    <div class="expenses-button desktop-btn">
        <button @click="showModal = true">
            Добавить расход
        </button>
    </div>

  <!-- модальное окно -->
    <div v-if="showModal" class="modal-overlay" @click.self="showModal = false">
      <div class="modal">
        <h2>Добавить расход</h2>

        <!-- Поля формы -->
        <label>
          Статья расхода
          <input type="text" v-model="newExpense.name" placeholder="Введите название">
        </label>

        <label>
          Сумма
          <input type="text" v-model="newExpense.amount" placeholder="Введите сумму">
        </label>

        <label>
          Счет оплаты
          <select v-model="newExpense.account">
            <option disabled value="">Выберите счет</option>
            <option v-for="acc in accounts" :key="acc" :value="acc">
              {{ acc }}
            </option>
          </select>
        </label>

        <label>
          Дата оплаты
          <input type="date" v-model="newExpense.date">
        </label>

        <label>Тег</label>

        <div class="tags-row">
          <div
            v-for="tag in tags"
            :key="tag.name"
            class="tag-option"
            :class="{ active: newExpense.tag === tag.name }"
            @click="selectTag(tag)"
          >
            <span
              class="tag-color"
              :style="{ background: tag.color }"
            ></span>

            {{ tag.name }}
          </div>
        </div>
        <div class="new-tag">
          <h4>Новый тег</h4>

          <input
            type="text"
            v-model="newTag.name"
            placeholder="Название тега"
          >

          <input
            type="color"
            v-model="newTag.color"
          >

          <button @click="addTag">
            Создать тег
          </button>
        </div>
        <label>
          Описание
          <textarea v-model="newExpense.description"></textarea>
        </label>

        <!-- Кнопки -->
        <div class="modal-buttons">
          <button class="save-btn" @click="addExpense">Сохранить</button>
          <button class="cancel-btn" @click="showModal = false">Отменить</button>
        </div>
      </div>
    </div>
  </section>

</template>

<script>
import Header from './Header.vue'
import Nav from './Nav.vue'

export default {

  components:{
    Header,
    Nav
  },

  data(){
    return{
      showModal: false,

      // новый расход
      newExpense: {
        name: '',
        amount: '',
        date: '',
        tag: '',
        color: '',
        account: ''
      },

      // счета
      accounts: [
        'Наличные',
        'Карта Сбер',
        'Тинькофф',
        'PayPal'
      ],

      // теги
      tags: [
        { name: 'Обучение', color: '#EE6F6F' },
        { name: 'Подписки', color: '#A3BA63' },
        { name: 'Сервисы', color: '#5170AD' },
        { name: 'Маркетинг', color: '#A346B2' }
      ],

      // создание нового тега
      newTag: {
        name: '',
        color: '#3383FB'
      },
      expenses:[

        {
        name:"Обучение верстке",
        amount:"45 000 р.",
        date:"30.07.2024",
        tag:"Обучение",
        color:"#EE6F6F"
        },

        {
        name:"Тильда",
        amount:"3000 р.",
        date:"21.07.2024",
        tag:"Подписки",
        color:"#A3BA63"
        },

        {
        name:"Nolim",
        amount:"550 р.",
        date:"21.07.2024",
        tag:"Сервисы",
        color:"#5170AD"
        },

        {
        name:"Anex",
        amount:"550 р.",
        date:"21.07.2024",
        tag:"Маркетинг",
        color:"#A346B2"
        }

        ]

    }
  },
  methods: {
    addExpense() {
      if(this.newExpense.name && this.newExpense.amount) {

        this.expenses.push({
          ...this.newExpense
        });

        this.newExpense = {
          name:'',
          amount:'',
          date:'',
          tag:'',
          color:'',
          account:''
        };

        this.showModal = false;
      } else {
        alert('Введите название и сумму');
      }
    },

    selectTag(tag) {
      this.newExpense.tag = tag.name;
      this.newExpense.color = tag.color;
    },

    addTag() {
      if(!this.newTag.name) return;

      const tag = {
        name: this.newTag.name,
        color: this.newTag.color
      };

      this.tags.push(tag);

      // сразу выбираем новый тег
      this.selectTag(tag);

      // сброс
      this.newTag = {
        name: '',
        color: '#3383FB'
      };
    }
  }

}
</script>

<style scoped>
* {
  box-sizing: border-box;
}

/* Заголовок страницы */
.page-title {
  text-align: center;
  font-family: "Manrope";
  font-weight: 500;
  font-size: 30px;
  line-height: 140%;
  letter-spacing: 0%;
  color: #000000;
  margin-bottom: 30px;
  width: 100%;
}

/* Мобильная кнопка */
.mobile-add-btn {
  display: none;
  position: absolute;
  top: 10px;
  right: 10px;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: #3383FB;
  color: white;
  border: none;
  font-size: 24px;
  cursor: pointer;
  z-index: 10;
  box-shadow: 0 2px 8px rgba(0,0,0,0.2);
  align-items: center;
  justify-content: center;
}

.mobile-add-btn span {
  line-height: 1;
}

.modal textarea {
  width: 100%;
  min-height: 100px;
  margin-top: 5px;
  padding: 10px;
  border-radius: 6px;
  border: 1px solid #ccc;
  background: #F1F4FF;
  resize: vertical;
  font-family: 'Manrope';
  font-size: 14px;
}

.modal select {
  width: 100%;
  height: 40px;
  margin-top: 5px;
  padding: 0 10px;
  border-radius: 6px;
  border: 1px solid #ccc;
  background: #F1F4FF;
  font-family: 'Manrope';
  font-size: 14px;
}

.tags-row {
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
  margin-bottom: 15px;
}

.tag-option {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 10px;
  border-radius: 6px;
  background: #F1F4FF;
  cursor: pointer;
  border: 1px solid transparent;
}

.tag-option.active {
  border: 1px solid #3383FB;
}

.new-tag {
  margin-top: 15px;
  display: flex;
  gap: 10px;
  align-items: center;
  flex-wrap: wrap;
}

.new-tag h4 {
  width: 100%;
  margin: 0 0 5px 0;
  font-family: 'Manrope';
  font-size: 14px;
  color: rgba(0,0,0,0.7);
}

.new-tag input[type="text"] {
  flex: 1;
  height: 40px;
  padding: 0 10px;
  border-radius: 6px;
  border: 1px solid #ccc;
  background: #F1F4FF;
  min-width: 150px;
}

.new-tag input[type="color"] {
  width: 40px;
  height: 40px;
  padding: 0;
  border: none;
  background: none;
  cursor: pointer;
}

.new-tag input[type="color"]::-webkit-color-swatch-wrapper {
  padding: 0;
}

.new-tag input[type="color"]::-webkit-color-swatch {
  border-radius: 6px;
  border: 1px solid #ccc;
}

.new-tag button {
  height: 40px;
  background: #3383FB;
  color: white;
  border: none;
  padding: 0 16px;
  border-radius: 6px;
  cursor: pointer;
  white-space: nowrap;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0,0,0,0.3);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 100;
}

.modal {
  background: #fff;
  width: 90%;
  max-width: 755px;
  max-height: 90vh;
  overflow-y: auto;
  padding: 30px;
  border-radius: 15px;
  position: relative;
}

.modal h2 {
  font-family: 'Manrope';
  font-size: 20px;
  color: rgba(0,0,0,0.7);
  margin-bottom: 20px;
}

.modal label {
  display: block;
  margin-bottom: 15px;
  font-size: 14px;
  color: rgba(0,0,0,0.7);
}

.modal input {
  width: 100%;
  height: 40px;
  margin-top: 5px;
  padding: 0 10px;
  border-radius: 6px;
  border: 1px solid #ccc;
  background: #F1F4FF;
}

.modal-buttons {
  display: flex;
  gap: 15px;
  margin-top: 30px;
}

.save-btn {
  background: #3383FB;
  color: #fff;
  border: none;
  border-radius: 4px;
  height: 35px;
  padding: 0 20px;
  cursor: pointer;
}

.cancel-btn {
  background: #E9E9E9;
  color: #000;
  border: none;
  border-radius: 4px;
  height: 35px;
  padding: 0 20px;
  cursor: pointer;
}

.expenses-button {
  width: 100%;
  max-width: 805px;
  margin: 20px auto 0;
  filter: drop-shadow(2px 2px 6px rgba(0,0,0,0.08));
}

.expenses-button button {
  width: 100%;
  height: 40px;
  background: #3383FB;
  border: none;
  border-radius: 15px;
  font-family: Manrope;
  font-size: 14px;
  font-weight: 400;
  color: white;
  cursor: pointer;
}

.expenses-button button:hover {
  background: #2e74df;
}

.expenses-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 60px 20px;
  background: #F1F4FF;
  min-height: 100vh;
}

/* контейнер таблицы */
.expenses-container {
  width: 100%;
  max-width: 805px;
  background: white;
  border: 0.5px solid #BBBBBB;
  border-radius: 15px;
  overflow: hidden;
  position: relative;
}

/* grid таблицы */
.table-grid {
  display: grid;
  grid-template-columns: 3fr 1.2fr 1.4fr 1fr 40px;
  align-items: center;
}

/* header */
.table-header {
  height: 53px;
  background: #DBE1F6;
  padding: 0 18px;
  font-family: Manrope;
  font-size: 14px;
  font-weight: 500;
  color: rgba(0,0,0,0.7);
  border-bottom: 0.5px solid #BBBBBB;
}

/* строки */
.table-row {
  height: 50px;
  padding: 0 18px;
  font-family: Manrope;
  font-size: 14px;
  font-weight: 300;
  color: rgba(0,0,0,0.7);
  border-bottom: 0.5px solid #BBBBBB;
}

.table-row:last-child {
  border-bottom: none;
}

/* расход + чекбокс */
.name {
  display: flex;
  align-items: center;
  gap: 12px;
}

/* тег цвет */
.tag {
  display: flex;
  align-items: center;
  gap: 8px;
}

.tag-color {
  width: 16px;
  height: 16px;
  border-radius: 2px;
  flex-shrink: 0;
}

.tag-text {
  font-family: Manrope;
  font-size: 14px;
  font-weight: 300;
  color: rgba(0,0,0,0.7);
}

.tag span {
  display: block;
  width: 16px;
  height: 16px;
  border-radius: 2px;
}

/* иконка редактирования */
.edit img {
  width: 16px;
  cursor: pointer;
}

.edit-icon {
  width: 16px;
  height: 16px;
}

/* Адаптивность для мобильных устройств */
@media (max-width: 768px) {
  /* Скрываем Nav */
  .nav-component {
    display: none !important;
  }
  
  /* Адаптируем заголовок */
  .page-title {
    font-size: 24px;
    margin-bottom: 20px;
  }
  
  /* Адаптируем страницу */
  .expenses-page {
    padding: 20px 10px;
  }
  
  /* Показываем мобильную кнопку */
  .mobile-add-btn {
    display: flex;
  }
  
  /* Скрываем десктопную кнопку */
  .desktop-btn {
    display: none;
  }
  
  /* Адаптируем контейнер */
  .expenses-container {
    width: 100%;
    max-width: 100%;
    border-radius: 10px;
  }
  
  /* Адаптируем grid */
  .table-grid {
    grid-template-columns: 2fr 1fr 1.2fr 0.8fr 30px;
    gap: 5px;
  }
  
  /* Адаптируем header */
  .table-header {
    padding: 0 10px;
    font-size: 11px;
    height: 45px;
  }
  
  /* Адаптируем строки */
  .table-row {
    padding: 0 10px;
    font-size: 11px;
    height: 45px;
  }
  
  /* Уменьшаем чекбокс */
  .name input[type="checkbox"] {
    width: 16px;
    height: 16px;
  }
  
  .name {
    gap: 8px;
  }
  
  /* Адаптируем тег */
  .tag {
    gap: 4px;
  }
  
  .tag-color {
    width: 12px;
    height: 12px;
  }
  
  .tag-text {
    font-size: 11px;
  }
  
  /* Адаптируем иконку редактирования */
  .edit-icon {
    width: 14px;
    height: 14px;
  }
  
  /* Адаптируем модальное окно */
  .modal {
    width: 95%;
    padding: 20px;
    margin: 10px;
    max-height: 85vh;
  }
  
  .modal h2 {
    font-size: 18px;
  }
  
  .modal input,
  .modal select {
    height: 36px;
    font-size: 14px;
  }
  
  .new-tag {
    flex-direction: column;
    align-items: stretch;
  }
  
  .new-tag input[type="text"] {
    min-width: auto;
  }
}

@media (max-width: 480px) {
  /* Адаптируем заголовок */
  .page-title {
    font-size: 20px;
    margin-bottom: 15px;
  }
  
  .expenses-page {
    padding: 15px 5px;
  }
  
  /* Сильнее сжимаем grid */
  .table-grid {
    grid-template-columns: 1.5fr 0.8fr 1fr 0.7fr 25px;
    gap: 3px;
  }
  
  .table-header {
    padding: 0 8px;
    font-size: 10px;
    height: 40px;
  }
  
  .table-row {
    padding: 0 8px;
    font-size: 10px;
    height: 40px;
  }
  
  .name {
    gap: 5px;
  }
  
  .name input[type="checkbox"] {
    width: 14px;
    height: 14px;
  }
  
  .tag-color {
    width: 10px;
    height: 10px;
  }
  
  .tag-text {
    font-size: 10px;
  }
  
  .mobile-add-btn {
    width: 36px;
    height: 36px;
    font-size: 20px;
    top: 8px;
    right: 8px;
  }
  
  .modal {
    padding: 15px;
  }
  
  .modal-buttons {
    flex-direction: column;
  }
  
  .save-btn, 
  .cancel-btn {
    width: 100%;
  }
  
  .tags-row {
    gap: 5px;
  }
  
  .tag-option {
    padding: 4px 8px;
    font-size: 12px;
  }
}
</style>