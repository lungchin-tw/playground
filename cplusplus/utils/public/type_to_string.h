
/*
** Copyright CHEN, LUNG-CHIN. All Rights Reserved.
*/


#pragma once


#include <string>


template<bool value>
const std::string& bool_to_string()
{
    static const std::string Value = "false";
    return Value;
};

template<>
const std::string& bool_to_string<true>()
{
    static const std::string Value = "true";
    return Value;
}

std::string address_to_string( const void* );

