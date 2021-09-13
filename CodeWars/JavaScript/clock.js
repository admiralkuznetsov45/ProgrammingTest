function past(h, m, s){
    //#Happy Coding! ^_^
    let hh , hm , hs;
    
    hh = h * 3600000 ; 
    hm = m * 60000 ; 
    hs = s * 1000 ; 
    
    return hs + hm + hh;
  }