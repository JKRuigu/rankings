import requests
import bs4
import sys
import re

# getting schools from 
# http://www.education.go.ke/results/Index.aspx
# this is data from schools that existed in 2012.
# New schools do not, therefore, exist
# Some schools could've closed down
# I do not know the code for private registration

def get_school_index_numbers():
    # exploits the fact that search is very fuzzy and not limited.
    data = []
    for x in xrange(10):
        payload = {'txtSchool':str(x),
                'SelectOrd':'TotalMarks+DESC',
                'Submit':'GO'}
        cookies={'ASP.NET_SessionId':'dajp54zyl4birripmf4c0d55'}
        r = requests.post(
                'http://www.education.go.ke/results/searchschool.aspx',
                data=payload,
                cookies=cookies)

        assert r.status_code == 200 # fail otherwise

        soup = bs4.BeautifulSoup(r.text)
        tbl = soup.find('table',
                {'border':1,
                'align':'center'}
                )
        labels = [el.text for el in tbl.contents[0].find_all('td')]
        #print labels
        for line in tbl.contents[1:]:
            try:
                contents = tuple(el.text.replace(u'\xa0',u'') for el in line.find_all('td'))
                assert len(contents) == len(labels) # fail otherwise
                data.append(contents)
            except AttributeError:
                #sys.stdout.write(line) # caused by empty contents
                pass

    #print data
    #print "data size: ", len(data)
    return set(data) # remove duplicates


def get_school_grades(s_index):
    pass

if __name__=='__main__':
    # get school index numbers
    indexes =get_school_index_numbers()
    #print "index length: ", len(indexes)
    #print indexes
    for s_index, s_name in indexes:
        re.sub('\s+', ' ', s_name)
        print "%s,%s"%(s_index, s_name,) # name and index


    sys.exit(1)

    # for each school, get candidate data
    for school_index in indexes:
        get_school_grades(school_index)


