/*
** Copyright CHEN, LUNG-CHIN. All Rights Reserved.
*/


#pragma once


template<class TYPE>
struct IsPointerType
{
    enum { Value = 0 };
};

template<class TYPE>
struct IsPointerType<TYPE*>
{
    enum { Value = 1 };
};