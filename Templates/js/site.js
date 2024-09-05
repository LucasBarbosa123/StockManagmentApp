function ChangeArrow(e) {
    let container = e.target

    if (e.target.tagName.toLowerCase() == 'i') {
        container = container.parentElement
    }

    let arrow = container.lastElementChild

    if (arrow.classList.contains('fa-chevron-right')) {
        arrow.classList.remove('fa-chevron-right')
        arrow.classList.add('fa-chevron-down')
    }
    else {
        arrow.classList.remove('fa-chevron-down')
        arrow.classList.add('fa-chevron-right')
    }
}

function DateNormalizer(unormalizedDate) {
	let year = unormalizedDate.substring(0, 4)
	let month = unormalizedDate.substring(5, 7)
	let day = unormalizedDate.substring(8, 10)

	return day + '/' + month + '/' + year
}

async function Logout(event) {
    event.preventDefault()
    let success = await fetch('/api/logout', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        }
    })

    if (success.ok) {
        location.reload()
    }
}