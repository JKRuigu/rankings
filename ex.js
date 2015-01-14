var page = new WebPage(), testindex = 0, badCount=0, loadInProgress = false, done=false;
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

                var pageIndexNumber=document.getElementById("ctl00_cphMain_TabContainer1_Marks_LblIndex").innerHTML;
                var pageschoolIndexNumber=document.getElementById("ctl00_cphMain_TabContainer1_Marks_LblCode").innerHTML;
                var schoolName=document.getElementById("ctl00_cphMain_TabContainer1_Marks_txtSchool").value;
                var gender=document.getElementById("ctl00_cphMain_TabContainer1_Marks_txtGender").value;
                var engGrade=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl02_GRADE").value;
                var kisGrade=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl03_GRADE").value;
                var matGrade=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl04_GRADE").value;
                var sciGrade=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl05_GRADE").value;
                var ssrGrade=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl06_GRADE").value;


                //console.log("eng ", eng, "kis ", kis, "mat ", mat, "sci ", sci,"ssr ", ssr ,"marks ", marks);
                console.log(pageIndexNumber+','+gender+','+name+','+eng+','+ kis+','+ mat+','+ sci+','+ ssr+','+ marks+','+schoolName+','+pageschoolIndexNumber+','+engGrade+','+kisGrade+','+matGrade+','+sciGrade+','+ssrGrade);
		badCount =0; //reset bad Count
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

console.log('SCHOOL::::',schoolCode);

var candidateInProgress = false,
    candidate=1;

page.onError = function(){
    if (badCount > 5){
        phantom.exit();
    }else {
        badCount++;
    }
    return;
}

page.open("http://www.knec-portal.ac.ke/RESULTS/ResultKCPE.aspx");

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
        }, 10);
    }
}, 45);

