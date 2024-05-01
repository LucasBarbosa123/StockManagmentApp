async function Login() {
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

function addEnterListeners() {
    let formInputs = document.querySelectorAll('#loginForm input')
    let loginBtt = document.getElementById('loginBtt')

    formInputs.forEach((input) => {
        input.addEventListener('keyup', (e) => {
            if(e.key.toLowerCase() == 'enter') {
                loginBtt.click()
            }
        })
    })
}

addEnterListeners()