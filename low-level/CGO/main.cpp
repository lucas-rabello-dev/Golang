#include <iostream>
extern "C" {
    void HelloFromCPP() {
        std::cout << "Olá do C++!" << std::endl;
    }
}
