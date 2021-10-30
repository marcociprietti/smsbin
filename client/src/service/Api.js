export default class Api {
  static getSmsList() {
    const request = new Request('http://localhost:10013/sms');

    return fetch(
      request,
      {
        method: 'GET',
        mode: 'cors',
      },
    )
      .then((response) => response.json());
  }
}
