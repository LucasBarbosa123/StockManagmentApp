function ChangeArrow(e) {
    let container = e.target

    if (e.target.tagName.toLowerCase() == 'i') {
        container = container.parentElement
    }

    let arrow = container.lastElementChild

    if (arrow.classList.contains('fa-chevron-left')) {
        arrow.classList.remove('fa-chevron-left')
        arrow.classList.add('fa-chevron-down')
    }
    else {
        arrow.classList.remove('fa-chevron-down')
        arrow.classList.add('fa-chevron-left')
    }
}

function DateNormalizer(unormalizedDate) {
	let year = unormalizedDate.substring(0, 4)
	let month = unormalizedDate.substring(5, 7)
	let day = unormalizedDate.substring(8, 10)

	return day + '/' + month + '/' + year
}