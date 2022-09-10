import { LitElement, html, css} from 'lit-element';
import { picnic as picnic } from './styles.js';
import { initcss as initcss } from './init-css';

class HelloWorld extends LitElement {
  static styles = [picnic]

  render() {
    return html`
    <h1 class="title">👋 Hello World 🌍</h1>
    `;
  }
}

customElements.define('hello-world', HelloWorld);

class ServedByCapsule extends LitElement {
  static styles = [picnic]

  render() {
    return html`
    <h2 class="subtitle">Served with 💜 by Capsule 💊</h2>
    `;
  }
}

customElements.define('served-by-capsule', ServedByCapsule);

class LittleMessage extends LitElement {
  static styles = [picnic]

  render() {
    return html`
    <div>
      <h2 class="subtitle"> {{message}} </h2>
      <fieldset class="flex two">
        <label><input type="email" placeholder="Email"></label>
        <label><input type="password" placeholder="Password"></label>
      </fieldset>
      <textarea placeholder="Textarea"></textarea>
    </div>
    `;
  }
}

customElements.define('little-message', LittleMessage);

class MainApp extends LitElement {
  static styles = [picnic, initcss]

  render() {
    return html`
    <div class="fontbase paddingbase">
      <section class="flex">
        <article class="card">
          <div>
            <hello-world></hello-world>
            <served-by-capsule></served-by-capsule>
            <little-message></little-message>
            <button class='success'>Success</button>
          </div>
        </article> 
      </section> 
    </div>
    `;
  }
}

customElements.define('main-app', MainApp);

/*
body {
  font-family: Algerian;
}
*/