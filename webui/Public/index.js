let bookmarkList=[]

async function loadDB() { 
  fetch("http://localhost:8081/bookmarks")
    .then((res)=>{
      return res.json()
    })
    .then ((data)=>{
      bookmarkList = data.bookmarks
      showBookmarks(bookmarkList)
    })
    .catch((err)=>{
      console.log("Error occurred retrieving database: ", err)
    })
}

const formEl = document.getElementById("form-bookmark")
formEl.addEventListener("submit", submitBookmark)

async function submitBookmark(e) {
  e.preventDefault()

  const formData = new FormData(formEl)
  let data = {}

  formData.forEach((value, key)=>{
    data[key] = value
  })

  console.log("data = ", data)

  const fOptions = {
    method: "POST",
    cache: "no-cache",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  }

  fetch("http://localhost:8081/bookmarks", fOptions)
    .then((res)=>{
      if (res.ok) {
        return res.json()
      }
    })
  .then((data)=>{
    addBookmark(data)
  })
}

function showBookmarks(bookmarks) {
  const target = document.getElementById("bookmark-list")

  target.innerHTML = ""

  for (let i=0; i < bookmarks.length; i++) {
    target.innerHTML += htmlBookmark(bookmarks[i], i) 
  }
}

const deleteBookmark=(loopId, dbId)=> {
  
  console.log("clickon: ", loopId)

  const ele = document.getElementById(loopId)

  while (ele.firstChild) {
    ele.removeChild(ele.firstChild)
  }

  fetch(`http://localhost:8081/bookmarks/${dbId}`, {method:"delete"})
    .then((res)=>{
      if (!res.ok) {
        console.log("expected ok, received: ", res.status)
      }
    })
    .catch((err)=>{
      console.log("error deleting bookmark; ", err)
    })

}

const addBookmark=(addedBookmark)=>{
  bookmarkList.push(addedBookmark)
  showBookmarks(bookmarkList)
}

const htmlBookmark=(bookmark, i=1)=>{
  
  const html=
   `<div id=${i} class="bookmark">
    <button type="button" class="button delete" onclick="deleteBookmark(${i},${bookmark.rowid})" >X</button>
    <h3><a href=${bookmark.url} target="_blank">${bookmark.url}</a></h3>
    -- ${bookmark.description}
    </div>`

  return html
}
