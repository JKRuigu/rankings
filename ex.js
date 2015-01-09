var page = new WebPage(), testindex = 0, loadInProgress = false, done=false;
var system = require('system');


page.onConsoleMessage = function(msg) {
  console.log(msg);
};

page.onLoadStarted = function() {
  loadInProgress = true;
  //console.log("load started");
};

page.onLoadFinished = function() {
  loadInProgress = false;
  //console.log("load finished");
};


function get_steps(index_number){
    //console.log('done is ', done);
    var steps = [
        function() {
            // load page TODO run this once
            page.open("http://www.knec-portal.ac.ke/RESULTS/ResultKCPE.aspx");
        },
        function() {
            //Enter index number and submit
            page.evaluate(function(index_number) {
                document.getElementById("ctl00_cphMain_TabContainer1_Marks_txtIndex").value=index_number;
                document.getElementById("ctl00_cphMain_TabContainer1_Marks_btnFind").click();
                return;
            }, index_number);
        }, 
        function() {
            //Get Marks
            page.evaluate(function(indexNumber) {
                var marks = document.getElementById("ctl00_cphMain_TabContainer1_Marks_txtTotal").value;
                var name = document.getElementById("ctl00_cphMain_TabContainer1_Marks_txtName").value;
                var eng=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl02_MKS").value;
                var kis=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl03_MKS").value;
                var mat=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl04_MKS").value;
                var sci=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl05_MKS").value;
                var ssr=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl06_MKS").value;

                //console.log("eng ", eng, "kis ", kis, "mat ", mat, "sci ", sci,"ssr ", ssr ,"marks ", marks);
                console.log(indexNumber+','+name+','+eng+','+ kis+','+ mat+','+ sci+','+ ssr+','+ marks);
                return;
            }, index_number);
        }
    ];

    return steps;
}


function padToThree(number) {
  number = ("000"+number).slice(-3);
  return number;
}


if (system.args.length === 1) {
    console.log('Try to pass some args when invoking this script!');
    phantom.exit();
} 

var schoolCode = system.args[1];

var candidateInProgress = false,
    candidate=1;

page.onError = function(){
    phantom.exit();
}

candidateInterval = setInterval(function(){
    if(!candidateInProgress){
        candidateInProgress = true;
        indexNumber = schoolCode + padToThree(candidate);
        //console.log('indexNumber ', indexNumber);
        var steps = get_steps(indexNumber); // sample school code and index number
        testindex=0;

        interval = setInterval(function() {
          if (!loadInProgress && typeof steps[testindex] == "function") {
            //console.log("step " + (testindex + 1));
              steps[testindex]();
              testindex++;
          }
          if (typeof steps[testindex] != "function") {
            //console.log("test complete!");
            //phantom.exit();
            //console.log('done with candidate ',indexNumber);
            clearInterval(interval);
            candidate++;
            candidateInProgress = false;
          }
        }, 50);
    }
}, 1000);


