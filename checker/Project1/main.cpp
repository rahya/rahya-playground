#include <cstdlib>



#include <set>
#include <vector>
#include <string>
#include <iostream>
#include <map>
#include <iomanip>
#include <sstream>


class AVLRE_Route {


public :
	AVLRE_Route(const std::string& iString): _string(iString) {}


	std::string _string;

	bool operator==(const AVLRE_Route& iOther) { return iOther._string == _string; }
	bool operator< (const AVLRE_Route& iOther) { return iOther._string <  _string; }

	

};

std::ostream& operator<<(std::ostream& iOut, const AVLRE_Route& iRoute) {
	iOut << iRoute._string;
	return iOut;
}



class CheckPrinter {

	struct test {
		int64_t testIndex;
		std::string result;
	};
	struct CompareTest
	{
		bool operator () (const test& lhs, const test& rhs) const { return lhs.testIndex < rhs.testIndex; }
		bool operator () (const test& lhs, const int64_t& rhs) const { return lhs.testIndex < rhs; }
		bool operator () (const int64_t& lhs, const test& rhs) const { return lhs < rhs.testIndex; }
	};

	/*
	struct TestInfo {
		const char* name;
	};

	struct Item {
		std::string name;
		std::set<test, CompareTest> tests;
	};
	struct CompareItem
	{
		bool operator() (const std::string& lhs, const std::string& rhs) const { return lhs < rhs; }
		bool operator() (const Item& lhs, const Item& rhs) const { return lhs.name < rhs.name; }
	};
	*/


	std::map<std::string, std::map<int64_t, std::string>> _map;

	std::vector<const char*> _testInfo;


	public:
		friend int main();

		
		void addTestInfo(const char* iName) {
			_testInfo.push_back( iName );
		}

		typename std::vector<const char*>::iterator searchTest(const char* iTestName) {
			auto itTestInfo = std::find_if(_testInfo.begin(), _testInfo.end(), [&iTestName](const char* iTestInfo) { return iTestInfo == iTestName; });
			if (itTestInfo == _testInfo.end()) {
				addTestInfo(iTestName);
				itTestInfo = std::find_if(_testInfo.begin(), _testInfo.end(), [&iTestName](const char* iTestInfo) { return iTestInfo == iTestName; });
				std::cout << "Test " << iTestName << " added at slot " << std::distance(_testInfo.begin(), itTestInfo) << std::endl;
			}
			else {
				std::cout << "Test " << iTestName << " found at slot " << std::distance(_testInfo.begin(), itTestInfo) << std::endl;

			}
			return itTestInfo;
		}


		
		
		void addTest(const std::string& iItemName, const char* iTestName, const std::string& result) {
			auto itTestInfo = searchTest(iTestName);

			std::cout << "Add result " << result << " for test " << *itTestInfo << " to " << iItemName << std::endl;
			
			_map[iItemName].insert(std::pair(std::distance(_testInfo.begin(), itTestInfo), result));

		}


#define LOG_CHECK_PRINTER(VAR) std::cout << VAR << std::endl
		void printAllTest() {
			std::string str;
			for (auto& iItem : _testInfo) {
				if (str.size()) { str += ", "; }
				str += iItem;
			}
			LOG_CHECK_PRINTER("All test : [ " << str << " ] ");
		}


		void printHeader() {
			std::ostringstream str;
			str << std::setw(30) << "" << " | ";
			for (auto& aTest : _testInfo) {
				str << aTest << " | ";
			}
			LOG_CHECK_PRINTER(str.str());
		}

		void print() {
			printHeader();
			for (auto& iItem : _map) {
				std::ostringstream str;
				
				str << std::setw(30) << iItem.first << " | ";

				for (size_t i = 0; i < _testInfo.size(); i+=1) {
					str << std::setw(strlen(_testInfo[i])) << iItem.second[i].substr(0, strlen(_testInfo[i])) << " | ";
				}
				LOG_CHECK_PRINTER(str.str());
			}
		}
		


	private:


};







int main() {

	CheckPrinter cp;


	cp.searchTest("test0");
	cp.searchTest("test0");
	cp.searchTest("test1");


	AVLRE_Route a("Route A");
	AVLRE_Route b("Route B");

	cp.addTest(a._string, "test0", "true");
	cp.addTest(a._string, "test1", "false");

	cp.addTest(b._string, "test0", "longuuue");
	cp.addTest(b._string, "test2    ", "X");

	cp.printAllTest();

	cp.print();




	 
	return EXIT_SUCCESS;
}