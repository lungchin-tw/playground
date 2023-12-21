
/*
** Copyright CHEN, LUNG-CHIN. All Rights Reserved.
*/


#include "type_to_string.h"
#include <sstream>


std::string address_to_string( const void* value )
{
    std::stringstream ss;
    ss << value;
    return ss.str();
}