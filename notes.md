- fetch
  response.json()

- http message
  https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages

example

```javascript
//post
// data to be sent to the POST request
let _data = {
  title: "foo",
  body: "bar",
  userId:1
}


fetch(url, {
    method: 'POST', // *GET, POST, PUT, DELETE, etc.
    headers: {
      'Content-Type': 'application/json'
      // 'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: JSON.stringify(data) // body data type must match "Content-Type" header
  });
.then(response => response.json())
.then(json => console.log(json));
.catch(err => console.log(err));

```
