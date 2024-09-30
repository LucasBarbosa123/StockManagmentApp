function OpenStufCreator() {
    let updtBtt = document.getElementById('UpdateBtt')
    let crtBtt = document.getElementById('CreateBtt')

    updtBtt.style.display = 'none'
    crtBtt.style.display = 'block'
}

function CreateStuf() {

}

function EditStuf() {

}

function HideShowColor() {
    let colorContainer = document.getElementById('colorContainer')
    let hasColor = document.getElementById('hasColor')

    if (hasColor.checked) {
        colorContainer.style.display = 'block'
    } else {
        colorContainer.style.display = 'none'
    }

}

function OpenColorCreator() {
    //hide the stuff modal before while going to the colors one
    $('#CreateUpdtModal').modal('hide')

    let colorName = document.getElementById('ColorName')
    let colorRef = document.getElementById('ColorRef')

    colorRef.value = '#000000'
    colorName.value = ''
}

function CloseColorModal() {
    //after we create the color we return to the stuff modal
    $('#CreateColorModal').modal('hide')
    $('#CreateUpdtModal').modal('show')
}

async function CreateColor() {
    let name = document.getElementById('ColorName')
    let ref = document.getElementById('ColorRef')
    let colorInfo = {
        Name: name.value,
        Ref: ref.value
    }

    let success = await CallColorCreator(colorInfo)
    if (!success) {
        alert('Algo de errado aconteceu ao criar a cor.')
        return
    }

    CloseColorModal()
}

async function CallColorCreator(colorInfo) {
    return new Promise((resolve, reject) => {
        fetch('/api/create-color', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			  },
			body: JSON.stringify(colorInfo)
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