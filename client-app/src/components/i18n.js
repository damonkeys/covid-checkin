import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';

import LanguageDetector from 'i18next-browser-languagedetector';
import XHR from 'i18next-xhr-backend';

// not like to use this?
// have a look at the Quick start guide 
// for passing in lng and translations on init

i18n
  .use(XHR)
  .use(LanguageDetector)
  .use(initReactI18next) // if not using I18nextProvider
  .init({
    backend: {
      loadPath: '/locales/{{lng}}/{{ns}}.json'
    },
    ns: ['common'],
    defaultNS: 'common',
    languages: ['en', 'de'],
    fallbackLng: 'en',
    debug: false,
    load: 'currentOnly',
    interpolation: {
      escapeValue: false, // not needed for react!!
    },
    detection: {
      // order and from where user language should be detected
      order: ['cookie', 'localStorage', 'navigator'],
      lookupCookie: 'i18n',
      lookupLocalStorage: 'i18nLng',

      // cache user language on
      caches: ['localStorage', 'cookie'],
      excludeCacheFor: ['cimode'], // languages to not persist (cookie, localStorage)

      // optional expire and domain for set cookie
      cookieMinutes: 10,
      cookieDomain: 'chckr.de'
    },
    // react i18next special options (optional)
    react: {
      wait: true,
      bindI18n: 'languageChanged loaded',
      bindStore: 'added removed',
      nsMode: 'default'
    }
  });


export default i18n;
