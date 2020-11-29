import firebase from 'firebase/app';
import "firebase/auth";

const firebaseConfig = {
  apiKey: "AIzaSyB1VefTFvO1z6oQUMR8p5uTD0rgKQQC58M",
  authDomain: "cut-match.firebaseapp.com",
  databaseURL: "https://cut-match.firebaseio.com",
  projectId: "cut-match",
  storageBucket: "cut-match.appspot.com",
  messagingSenderId: "236155511209",
  appId: "1:236155511209:web:08eeecbc2ef24d3be524b8",
  measurementId: "G-LF8F1CVW57"
};

export const myFirebase = firebase.initializeApp(firebaseConfig);
