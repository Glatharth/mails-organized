const fs = require('fs');
const Buffer = require('buffer');
const file_open = process.argv[2];
const file_save = process.argv[3];

if (!file_open) throw console.log("\nparâmetro file_open não declarado!");
if (!file_save) throw console.log("\nparâmetro file_save não declarado!");

fs.readFile(`./${file_open}`, (err, data) => {
    if (err) throw err;
    const regex = /^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,4}/gm
    var emails = data.toString().match(regex);
    var results = "Emails:\n";

    emails.forEach(i => {
        results += i + '\n';
    });

    fs.writeFile(`./${file_save}`, results, (err) => {
        if (err) throw err;
        console.log(`./${file_save}: ` + 'Emails salvos com sucesso.');
    });
});