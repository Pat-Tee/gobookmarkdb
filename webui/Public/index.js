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
  
  const ele = document.getElementById(loopId)

  fetch(`http://localhost:8081/bookmarks/${dbId}`, {method:"delete"})
    .then((res)=>{
      if (!res.ok) {
        console.log("expected ok, received: ", res.status)
      }
      while (ele.firstChild) {
        ele.removeChild(ele.firstChild)
      }
      bookmarkList.splice(loopId, 1)
    })
    .catch((err)=>{
      console.log("error deleting bookmark; ", err)
    })
}

const addBookmark=(addedBookmark)=>{
  bookmarkList.push(addedBookmark)
  showBookmarks(bookmarkList)
  document.getElementById("form-bookmark").reset()
}

const editBookmark=(loopId, dbId)=>{

  const bmEl = document.getElementById(loopId)

  bmEl.outerHTML = htmlEditBookmark(bookmarkList[loopId], loopId)
}

const saveBookmark=(bookmark)=>{

  console.log("save bookmark: ", bookmarkList[bookmark])
  const formData = new FormData(document.getElementById("editedBM"))
  let data = {}

  formData.forEach((value, key)=>{
    data[key] = value
  })

  const fOptions = {
    method: "PUT",
    cache: "no-cache",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  }

  fetch(`http://localhost:8081/bookmarks/${bookmarkList[bookmark].rowid}`, fOptions)
    .then((res)=>{
      if (res.ok) {
        return res.json()
      }
    })
  .then((data)=>{
     document.getElementById(bookmark).outerHTML = htmlBookmark(data, bookmark)
  })
}

const htmlBookmark=(bookmark, i=1)=>{
  
  const html=
   `<div id=${i} class="bookmark">
      <div class="button-group">
        <button type="button" class="button delete" onclick="deleteBookmark(${i},${bookmark.rowid})" >X</button>
        <button type="button" class="button edit" onclick="editBookmark(${i},${bookmark.rowid})" >E</button>
      </div>
      <div id="bmURL">
        <h3><a href=${bookmark.url} target="_blank">${bookmark.url}</a></h3>
      </div>
      <p id="bmDesc">${bookmark.description}</p>
    </div>`

  return html
}

const htmlEditBookmark=(bookmark, i)=>{

  const html=
   `<div id=${i} class="bookmark">
      <div class="button-group">
      <button type="button" class="button save" onclick="saveBookmark(${i},${bookmark.rowid})">Save</button>
      </div>
      <form id="editedBM">
        <input type="text" name="URL" value="${bookmarkList[i].url}">
        <input type="text" name="Description" value="${bookmarkList[i].description}">
      </form>
    </div>`

  return html
}

