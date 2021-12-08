import firebase from 'firebase/compat/app'
import 'firebase/auth'

if (!firebase.apps.length) {
  firebase.initializeApp({
    apiKey: 'AIzaSyCkACSnQluxeWGrLVRLUae4lskVV8li9Nw',
    authDomain: 'ponica-f7a81.firebaseapp.com',
    projectId: 'ponica-f7a81',
    storageBucket: 'ponica-f7a81.appspot.com',
    messagingSenderId: '1068883566081',
    appId: '1:1068883566081:web:25b9815f497dcacf4d6a00',
  })
}

export default firebase
