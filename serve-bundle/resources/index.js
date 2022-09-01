import { LitElement, html, css} from 'lit-element';
import { myStyle } from './style.js';

console.log("👋 Hello World 🌍")

class HelloWorld extends LitElement {
  static styles = [myStyle]

  render() {
    return html`
    <h1 class="title">👋 Hello World 🌍</h1>
    `;
  }
}

customElements.define('hello-world', HelloWorld);

class ServedByCapsule extends LitElement {
  static styles = [myStyle]

  render() {
    return html`
    <h2 class="subtitle">Served with 💜 by Capsule 💊</h2>
    `;
  }
}

customElements.define('served-by-capsule', ServedByCapsule);

class LittleMessage extends LitElement {
  static styles = [myStyle]

  render() {
    return html`
    <h2 class="subtitle"> {{message}} </h2>
    `;
  }
}

customElements.define('little-message', LittleMessage);

class MainApp extends LitElement {
  static styles = [myStyle]

  render() {
    return html`
    <section class="container">
      <div>
        <hello-world></hello-world>
        <served-by-capsule></served-by-capsule>
        <little-message></little-message>
      </div>
    </section>    
    `;
  }
}

customElements.define('main-app', MainApp);
