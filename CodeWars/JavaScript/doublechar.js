
//cara yang jauh lebih sederhana
function double(str) {
    return str.split("").map(i => String(i).repeat(2)).join("")
  }
  
  console.log(double("Baju")) // "BBaajjuu"

//cara awal
function doubleChar(str){

    //melakukan inisiasi hasil
    let hasil =""

    //melakukan iterasi agar char nya double 
    for (i=0;i<str.length;i++){
            hasil += str[i]+ str[i] 
    }
    console.log(hasil)
}

doubleChar("Baju")


