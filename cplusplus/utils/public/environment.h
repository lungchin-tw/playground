/*
** Copyright CHEN, LUNG-CHIN. All Rights Reserved.
*/


#pragma once 


class Environment
{
public:

	static Environment& Get();

	
	Environment( const Environment& ) = delete;
	Environment( Environment&& ) = delete;

	Environment& operator=( const Environment& ) = delete;
	Environment& operator=( Environment&& ) = delete;

	void Print();

private:
	Environment(){}
	~Environment(){}
};
