var page = new WebPage(), testindex = 0, loadInProgress = false;

page.onConsoleMessage = function(msg) {
  console.log(msg);
};

page.onLoadStarted = function() {
  loadInProgress = true;
  console.log("load started");
};

page.onLoadFinished = function() {
  loadInProgress = false;
  console.log("load finished");
};


function get_steps(index_number){
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
            page.evaluate(function() {
                var marks = document.getElementById("ctl00_cphMain_TabContainer1_Marks_txtTotal").value;
                var eng=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl02_MKS").value;
                var kis=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl03_MKS").value;
                var mat=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl04_MKS").value;
                var sci=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl05_MKS").value;
                var ssr=document.getElementById("ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl06_MKS").value;

                console.log("eng ", eng, "kis ", kis, "mat ", mat, "sci ", sci,"ssr ", ssr ,"marks ", marks);
                return;
            });
        }, 
        function() {
            page.evaluate(function() {
                console.log('done');
            });
        }
    ];

    return steps;
}

steps = get_steps("40735220001"); // sample school code and index number

interval = setInterval(function() {
  if (!loadInProgress && typeof steps[testindex] == "function") {
    //console.log("step " + (testindex + 1));
    steps[testindex]();
    testindex++;
  }
  if (typeof steps[testindex] != "function") {
    console.log("test complete!");
    phantom.exit();
  }
}, 50);
