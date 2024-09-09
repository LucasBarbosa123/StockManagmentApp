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

    console.log(hasColor)
    console.log(hasColor.checked)

    if (hasColor.checked) {
        colorContainer.style.display = 'block'
    } else {
        colorContainer.style.display = 'none'
    }

}