/*
** Copyright CHEN, LUNG-CHIN. All Rights Reserved.
*/


#include "environment.h"
#include "type_define.h"

#include <iostream>


Environment& Environment::Get()
{
	static Environment Instance;
	return Instance;
}

void Environment::Print()
{
	std::cout << "========== Environment::Print ===============" << std::endl;
	std::cout << "\t Sizeof( int )   : " << sizeof(int) << std::endl;
	std::cout << "\t Sizeof( int* )  : " << sizeof(int_ptr) << std::endl;
	std::cout << "\t Sizeof( int** ) : " << sizeof(int_ptr_ptr) << std::endl;
	std::cout << "=============================================" << std::endl;
}

