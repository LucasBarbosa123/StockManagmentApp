let warehouseBeenEdited = ''

function FormatWearHousesTbl() {
    $('#WearHousesTbl').DataTable()
}

function NormalizeTblDates() {
	let dateContainers = document.querySelectorAll('#WearHousesTbl .dateTD')

	dateContainers.forEach(container => {
		let normalizedDate = DateNormalizer(container.innerText)

		container.innerText = normalizedDate
	})
}

function OpenWearhouseCreator() {
    let UpdateBtt = document.getElementById('UpdateBtt')
    let CreateBtt = document.getElementById('CreateBtt')
    let NameInput = document.getElementById('Name')

    NameInput.value = ''
    CreateBtt.style.display = 'inline-block'
	UpdateBtt.style.display = 'none'
}

async function OpenWearhouseEditor(id) {
    let UpdateBtt = document.getElementById('UpdateBtt')
    let CreateBtt = document.getElementById('CreateBtt')
    let NameInput = document.getElementById('Name')

    warehouseBeenEdited = id

    NameInput.value = await GetWarehouseName(id)
    CreateBtt.style.display = 'none'
	UpdateBtt.style.display = 'inline-block'
}

async function GetWarehouseName(id) {
    return new Promise((resolve, reject) => {
        fetch('/api/warehouse-name/' + id, {
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

async function DeleteWearhouse(id) {
    let success = await CallWearhouseDeleter(id)

    if (success == true) {
        location.reload()
        return
    }
    alert(success)
}

async function CallWearhouseDeleter(id) {
    return new Promise((resolve, reject) => {
        fetch('/api/delete-warehouse', {
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

async function CreateWearhouse() {
    let name = document.getElementById('Name').value
    let success = await CallWearhouseCreator(name)

    if (success == true) {
        location.reload()
        return
    }

    alert('Algo inesperado ocourreu ao criar o armazem.')
}

async function CallWearhouseCreator(name) {
    return new Promise((resolve, reject) => {
        fetch('/api/create-wearhouse', {
            method: 'POST',
    	    headers: {
    	      'Content-Type': 'application/json'
    	    },
		    body: JSON.stringify(name)
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

async function EditWearhouse() {
    let name = document.getElementById('Name').value
    let warehouseInfo = {
        id: warehouseBeenEdited,
        name: name
    }

    console.log(warehouseInfo)

    let success = await CallWarehouseEditor(warehouseInfo)

    if (success == true) {
        location.reload()
        return
    }

    alert(success)
}

async function CallWarehouseEditor(warehouseInfo) {
    return new Promise((resolve, reject) => {
        fetch('/api/edit-warehouse', {
            method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			  },
			body: JSON.stringify(warehouseInfo)
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

NormalizeTblDates()
FormatWearHousesTbl()