import { createApp } from "vue";
import App from "./App.vue";
import { createI18n } from "vue-i18n"; 

const i18n = createI18n({
  locale: "en",
  messages: {
    en: {
      welcome: "Welcome to our application",
      simpleSpa: "Simple and stylish Vue.js SPA",
      footerAbout: "2025 Super application. All rights are protected.",
      card1Title: "Card 1",
      card2Title: "Card 2",
      card3Title: "Card 3",
      card1Text: "This is the text for card 1",
      card2Text: "This is the text for card 2",
      card3Text: "This is the text for card 3",
    },
    ru: {
      welcome: "Добро пожаловать в наше приложение",
      simpleSpa: "Простое и стильное SPA на Vue.js",
      footerAbout: "2025 Крутое приложение. Все права защищены.",
      card1Title: "Карточка 1",
      card2Title: "Карточка 2",
      card3Title: "Карточка 3",
      card1Text: "Это пример текста для первой карточки. Здесь может быть любая информация.",
      card2Text: "Это пример текста для второй карточки. Здесь может быть любая информация.",
      card3Text: "Это пример текста для третьей карточки. Здесь может быть любая информация.",
    },
  },
});

const app = createApp(App);
app.use(i18n);
app.mount("#app");