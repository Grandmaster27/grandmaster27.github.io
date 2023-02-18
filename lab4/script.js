let laptops = [
    {
        "model": "Lenovo Ideapad Gaming 3",
        "year": 2021,
        "price": 430000
    },
    {
        "model": "Lenovo Legion 5",
        "year": 2019,
        "price": 600000
    },
    {
        "model": "ASUS Slim Book",
        "year": 2023,
        "price": 380000
    }
];

function create_laptop(model, year, price){
    let new_laptop = {
        "model": model,
        "year": year,
        "price": price
    }
    return new_laptop;
}

laptops.push(create_laptop("Acer Extensa EX215-52", 2020, 260000));

function avg(laptops) {
    let summa = 0;
    for (let i = 0; i < laptops.length; i++) {
        summa += laptops[i]['price'];
    }
    return summa / laptops.length;
}





for (let i = 0; i < laptops.length; i++){
    document.write("<h2>Laptop model: " + laptops[i]["model"] + "</h2>")
    document.write("<h2>Year of release: " + laptops[i]["year"] + "</h2>")
    document.writeln("<h2>Price: " + laptops[i]["price"] + "</h2>")
}

document.write("<h2>Average price:" + avg(laptops) + "</h2>")