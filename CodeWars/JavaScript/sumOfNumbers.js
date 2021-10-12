function getSum( a,b )
{
  //sum = 0 , biasakan
  let sum = 0;

///// Solusi Yang Pertama 
// kondisi apabila a < b , maka a diloop sampai dgn b
//   if (a<b){
//     for (i = a ; i <= b ; i++){
//       sum += i; 
//     }
//   } else if (a>b){
//     for (i = b ; i <= a ; i++){
//       sum += i; 
//     }
//   } else if (a ==b) {
//     sum = a;
//   }
  
//   return sum;
// }

//// Solusi yang kedua 
if (a >b) {
  [a,b] = [b,a];
}

while (a<=b){
  sum += a;
  a++;
}

return sum;

}

console.log(getSum(1,2))