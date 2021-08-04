const get = document.getElementById("get")
const post = document.getElementById("post")
const content = document.querySelector(".content")
const del = document.getElementById("delete")



get.addEventListener("click", ()=>{
    fetch('/ping')
    .then((response) =>  response.json())
    .then((data)=> {
        console.log(data)
        output= "<h2>student list </h2>"
        data.forEach((student)=>{
            output += `
            <ul class="student">
                <li>name: ${student.name}</li>
                <li>age: ${student.age}</li>
                <li>hobby: ${student.hobby}</li>
            </ul>`
        })
        content.innerHTML = output
       
    })
    .catch((err)=>console.log(err))

})


post.addEventListener("click",(e)=>{

    e.preventDefault();

    let myForm = document.getElementById('myform');
    let formData = new FormData(myForm);

    fetch("/ping",{
        method: 'POST',
        body:formData
    })
    .then(response => response.json())
    .then(result => {
     console.log('Success:', result);
    })
    .catch(error => {
     console.log(error);
    });

})

del.addEventListener("click", ()=>{
    let name = document.getElementById("name")
    console.log(name.value)
    myform = new FormData()
    myform.append('name', name.value)

    fetch("/ping",{
        method:"DELETE",
        body: myform
    })
})