async function ChangeInfo(id) {
    let fName = document.getElementById('firstName')
    let lName = document.getElementById('lastName')

    if (fName == '' || lName == '') {
        alert('É necessário preencher o primeiro e segundo nome.')
    }

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

async function ChangePassword(id) {
    let oldPass = document.getElementById('oldPass').value
    let newPass = document.getElementById('newPass').value
    let newPass2 = document.getElementById('newPass2').value

    if (oldPass == '' || newPass == '' || newPass2 == '') {
        alert('Todos os campos precisam de estar prênchidos.')
        return
    }

    if (newPass != newPass2) {
        alert('Os campos de password nova têm que ser iguais.')
        return
    }

    let info = {
        oldPass: oldPass,
        newPass: newPass
    }

    let success = await CallPasswordChanger(id, info)
    if (!success) {
        alert('Algo de errado aconteceu ao tentar mudar a sua password.')
        return
    }
}

async function CallPasswordChanger(id, info) {
    return new Priomise((resolve, reject) => {
        fetch(`/api/change-user-pass/${id}`, {
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