async function ChangeInfo(id) {
    let fName = document.getElementById('firstName')
    let lName = document.getElementById('lastName')

    let info = {
        firstName: fName.value,
        lastName: lName.value,
    }

    let success = await CallInfoChanger(id, info)
    if (success == false) {
        alert('There was a problem updating the user info.')
        return
    }

    location.reload()
}

async function CallInfoChanger(id, info) {
    return new Promise((resolve, reject) => {
        fetch(`/api/change-user-info/${id}`, {
            method: 'POST',
            headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(info)
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
			resolve(false)
		})
    })
}

async function ChangeActualPassword() {
    console.log('')
}

async function CallPasswordChanger() {
    return new Priomise((resolve, reject) => {
        fetch('', {
            
        })
    })
}