
/*
** Copyright SOFTSTAR ENTERTAINMENT INC. All Rights Reserved.
*/


#pragma once


#if defined(_MSC_VER)
    #define FUNCTION_NAME __FUNCSIG__
#elif defined(__GNUC__) || defined(__clang__)
    #define FUNCTION_NAME __PRETTY_FUNCTION__
#else
    #define FUNCTION_NAME __FUNCTION__
#endif
