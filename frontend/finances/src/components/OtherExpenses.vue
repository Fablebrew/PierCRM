<template>
  <Header />
  <Nav page-title="Прочие расходы"/>

  <section class="expenses-page">

    <div class="expenses-container">

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
          {{ expense.name }}
        </div>

        <div>{{ expense.amount }}</div>
        <div>{{ expense.date }}</div>

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
          <img src="@/assets/img/redact.svg">
        </div>
      </div>
    </div>
    <div class="expenses-button">
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
.modal textarea {
  width: 100%;
  min-height: 100px;
  margin-top: 5px;
  padding: 10px;
  border-radius: 6px;
  border: 1px solid #ccc;
  background: #F1F4FF;
  resize: vertical; /* можно растягивать только по вертикали */
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
}

/* большой инпут */
.new-tag input[type="text"] {
  flex: 1;
  height: 40px;
  padding: 0 10px;
  border-radius: 6px;
  border: 1px solid #ccc;
  background: #F1F4FF;
}

/* маленький выбор цвета */
.new-tag input[type="color"] {
  width: 40px;
  height: 40px;
  padding: 0;
  border: none;
  background: none;
  cursor: pointer;
}

/* убираем стандартные отступы у color */
.new-tag input[type="color"]::-webkit-color-swatch-wrapper {
  padding: 0;
}

.new-tag input[type="color"]::-webkit-color-swatch {
  border-radius: 6px;
  border: 1px solid #ccc;
}

/* кнопка */
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
  width: 755px;
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
.expenses-button{

  width:805px;
  margin:20px auto 0;

  filter: drop-shadow(2px 2px 6px rgba(0,0,0,0.08));

}

.expenses-button button{

  width:100%;
  height:40px;

  background:#3383FB;

  border:none;
  border-radius:15px;

  font-family:Manrope;
  font-size:14px;
  font-weight:400;

  color:white;

  cursor:pointer;

}

.expenses-button button:hover{

  background:#2e74df;

}
.expenses-page{

  display:flex;
  flex-direction: column; /* <- добавить */
  align-items: center; 
  padding:60px 20px;
  background:#F1F4FF;
  min-height:100vh;

}

/* контейнер таблицы */

.expenses-container{

  width:805px;
  height:298px;

  background:white;
  border:0.5px solid #BBBBBB;
  border-radius:15px;

  overflow-y:auto;
}

/* grid таблицы */

.table-grid{

  display:grid;

  grid-template-columns:

  3fr
  1.2fr
  1.4fr
  1fr
  40px;

  align-items:center;

}

/* header */

.table-header{

  height:53px;

  background:#DBE1F6;

  padding:0 18px;

  font-family:Manrope;
  font-size:14px;
  font-weight:500;

  color:rgba(0,0,0,0.7);

  border-bottom:0.5px solid #BBBBBB;

}

/* строки */

.table-row{

  height:50px;
  padding:0 18px;

  font-family:Manrope;
  font-size:14px;
  font-weight:300;

  color:rgba(0,0,0,0.7);

  border-bottom:0.5px solid #BBBBBB;

}

.table-row:last-child{
  border-bottom:none;
}

/* расход + чекбокс */

.name{

  display:flex;
  align-items:center;
  gap:12px;

}

/* тег цвет */
.tag{

  display:flex;
  align-items:center;
  gap:8px;

}
.tag-color{

  width:16px;
  height:16px;

  border-radius:2px;

}
.tag-text{

  font-family:Manrope;
  font-size:14px;
  font-weight:300;

  color:rgba(0,0,0,0.7);

}

.tag span{

  display:block;

  width:16px;
  height:16px;

  border-radius:2px;

}

/* иконка редактирования */

.edit img{

  width:16px;
  cursor:pointer;

}



</style>