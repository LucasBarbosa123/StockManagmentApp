async function Login(event) {
    event.preventDefault()
    
    let email = document.getElementById('Email').value
    let pass = document.getElementById('Pass').value

    let loginInfo = {
        email: email,
        password: pass
    }

    let res = await CallLogin(loginInfo)

    if (res == true) {
        window.location.href = '/'
        return
    }

    alert(res)
}

async function CallLogin(loginInfo) {
    return new Promise((resolve, reject) => {
        fetch('/api/login', {
            method: 'POST',
            headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(loginInfo)
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