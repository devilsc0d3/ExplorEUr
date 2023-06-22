const sendData = async (txt) => {
    await fetch('http://localhost:8080/info', {
        method: 'POST',
        body: new URLSearchParams({
            categoryName: txt,
        }),
    });
};

function addCategory() {
    let button = document.querySelector("[name='submitButton']")
    button.addEventListener("click", (event) => {
        console.log("test")

        event.preventDefault();

        let form = document.querySelector('form');
        let categoryName = form.elements.categoryName.value;
        console.log(categoryName)

        let newCategoryDiv = document.createElement('div');
        newCategoryDiv.className = 'cat';
        let newCategoryParagraph = document.createElement('p');
        newCategoryParagraph.className = 'txt';
        newCategoryParagraph.textContent = categoryName;
        newCategoryDiv.appendChild(newCategoryParagraph);
        document.body.appendChild(newCategoryDiv);

        sendData(categoryName).then(res => res.json());

        console.log(categoryName)
        form.reset();
    })
}

const aCategory = [...document.querySelectorAll('a')].filter((el) => el.dataset.name === 'category');
let j = 1;
for (let i = 0; i < aCategory.length; i++) {
    aCategory[i].href = `/${j}`;
    j++;
}

export function category() {
    addCategory();
}
