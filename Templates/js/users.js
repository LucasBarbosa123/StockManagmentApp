let userBeenEdited = ''

function FormatUserTbl() {
    $('#UsersTbl').DataTable()
}

function NormalizeTblDates() {
	let dateContainers = document.querySelectorAll('#UsersTbl .dateTD')

	dateContainers.forEach(container => {
		let normalizedDate = DateNormalizer(container.innerText)

		container.innerText = normalizedDate
	})
}

async function GeneratePass() {
    let passInput = document.getElementById('Password')
    let pass = await GetGeneratedPass()

    passInput.value = pass
}

async function GetGeneratedPass() {
    return new Promise((resolve, reject) => {
        fetch('/api/generate-pass', {
          	method: 'GET',
          	headers: {
          	  'Content-Type': 'application/json'
          	}
        })
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok')
          }
          return response.json()
        })
        .then(data => {
          resolve(data)
        })
        .catch(error => {
          resolve("")
        })
    })
}

async function CreateUser() {
    let firstName = document.getElementById('FirstName').value
    let lastName = document.getElementById('LastName').value
    let email = document.getElementById('Email').value
    let pass = document.getElementById('Password').value

	let userObj = {
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Img: '',
		Password: pass
	}

	let result = await CallUserCreator(userObj)

	if (result != true) {
		alert(result)
		return
	}

	location.reload()
}

async function CallUserCreator(userObj) {
	return new Promise((resolve, reject) => {
      fetch('/api/create-user', {
        	method: 'POST',
        	headers: {
        	  'Content-Type': 'application/json'
        	},
			body: JSON.stringify(userObj)
        })
        .then(response => {
          if (!response.ok) {
			      throw response.json()
          }
          return response.json()
        })
        .then(data => {
          resolve(data)
        })
        .catch(error => {
          resolve(error)
        })
    })
}

async function DeleteUser(id) {
	let success = await CallUserDeletor(id)

	if (success != true) {
		alert(success)
		return
	}

	location.reload()
}

async function CallUserDeletor(id) {
  return new Promise((resolve, reject) => {
    	fetch('/api/delete-user', {
    	  method: 'POST',
    	  headers: {
    	    'Content-Type': 'application/json'
    	  },
		  body: JSON.stringify(id)
    	})
		.then(response => {
			if (!response.ok) {
				throw response.json()
			}
			return response.json()
		})
		.then(data => {
			resolve(data)
		})
		.catch(error => {
			resolve(error)
		})
  	})
}

async function OpenUserEditor(id) {
	let firstNameInput = document.getElementById('FirstName')
    let lastNameInput = document.getElementById('LastName')
    let emailInput = document.getElementById('Email')
	let passInputContainer = document.getElementById('PassInputContainer')
	let createBtt = document.getElementById('CreateBtt')
	let updateBtt = document.getElementById('UpdateBtt')

	try {
		let userInfo = await GetUserInfo(id)

		firstNameInput.value = userInfo.firstname
		lastNameInput.value = userInfo.lastname
		emailInput.value = userInfo.email

		passInputContainer.style.display = 'none'
		createBtt.style.display = 'none'
		updateBtt.style.display = 'inline-block'
	} catch (error) {
		alert(error)
		return
	}
}

async function GetUserInfo(id) {
	userBeenEdited = id

	return new Promise((resolve, reject) => {
		fetch('/api/user-info/' + userBeenEdited, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			}
		})
		.then(response => {
			if (!response.ok) {
				throw response.json()
			}
			return response.json()
		})
		.then(data => {
			resolve(data)
		})
		.catch(error => {
			reject(error)
		})
	})
}

function OpenUserCreator() {
	let firstNameInput = document.getElementById('FirstName')
    let lastNameInput = document.getElementById('LastName')
    let emailInput = document.getElementById('Email')
    let passInputContainer = document.getElementById('PassInputContainer')
	let createBtt = document.getElementById('CreateBtt')
	let updateBtt = document.getElementById('UpdateBtt')

	firstNameInput.value = ''
	lastNameInput.value = ''
	emailInput.value = ''

	passInputContainer.style.display = 'block'
	createBtt.style.display = 'inline-block'
	updateBtt.style.display = 'none'
}

async function EditUser() {
	let firstName = document.getElementById('FirstName').value
    let lastName = document.getElementById('LastName').value
    let email = document.getElementById('Email').value

	let userInfo = {
		id: userBeenEdited,
		firstname: firstName,
		lastname: lastName,
		email: email
	}
	
	let success = await CallUserEditor(userInfo)

	if (success != true) {
		alert(success)
		return
	}

	location.reload()
}

async function CallUserEditor(userInfo) {
	return new Promise((resolve, reject) => {
		fetch('/api/edit-user', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			  },
			body: JSON.stringify(userInfo)
		})
		.then(response => {
			if (!response.ok) {
				throw response.json()
			}
			return response.json()
		})
		.then(data => {
			resolve(data)
		})
		.catch(error => {
			resolve(error)
		})
	})
}

//functions that run imidiatly after loading
NormalizeTblDates()
FormatUserTbl()